package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/3almadmoon/ameni-assignment/config"
	pb "github.com/3almadmoon/ameni-assignment/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

type commandHolder interface {
	init([]string) error
	run() (interface{}, error)
	name() string
}

type cmdName string
type flagName string

const (
	add       cmdName = "add"
	update    cmdName = "update"
	deleteOne cmdName = "deleteOne"
	getAll    cmdName = "get-all"
	help      cmdName = "help"
)

const (
	title       flagName = "title"
	description flagName = "description"
	status      flagName = "status"
	hash        flagName = "hash"
)

const cliUsage = `[CLIENT CLI TOOL]:

USAGE:		<cmd> [--flag]  [arg]

Examples:
	 add --title [titleValue] --status [statusValue] --description [description] 

	 update -- hash [unique id] (required)

	 deleteOne -- hash [unique id] (required)

	 get-all (with no argument)			
`

type command struct {
	fs         *flag.FlagSet
	flagValues map[flagName]interface{}
}

type cmdFlag struct {
	name      flagName
	usage     string
	valueType string
}

type clientHandler struct {
	cancel context.CancelFunc
	close  func() error
	pb.TodoListServiceClient
}

// define flags name, usage and type for each command
var (
	addCmdFlags = []cmdFlag{
		{name: title, usage: "title of todo item", valueType: "string"},
		{name: description, usage: "description of todo item", valueType: "string"},
		{name: status, usage: "status of todo item", valueType: "int"},
	}
	updateCmdFlags = []cmdFlag{
		{name: hash, usage: "unique identifier of todo item", valueType: "string"},
		{name: status, usage: "status of todo item", valueType: "int"},
	}
	deleteCmdFlags = []cmdFlag{
		{name: hash, usage: "unique identifier of todo item", valueType: "string"},
	}
	getAllCmdFlags = []cmdFlag{
		{usage: "no argument expected"},
	}
)

// newCommand creates a command with its flagSet and map of flag values
func newCommand(cmdName cmdName, cmdFlags []cmdFlag) command {
	gc := command {
		fs:         flag.NewFlagSet(string(cmdName), flag.ExitOnError),
		flagValues: make(map[flagName]interface{}, 3),
	}
	for _, cmdFlag := range cmdFlags {
		switch cmdFlag.valueType {
		case "string":
			gc.flagValues[cmdFlag.name] = gc.fs.String(string(cmdFlag.name), "", cmdFlag.usage)
		case "int":
			gc.flagValues[cmdFlag.name] = gc.fs.Int(string(cmdFlag.name), 0, cmdFlag.usage)
		default:
			gc.fs.String(string(cmdFlag.name), "", cmdFlag.usage)
			return gc
		}
	}
	return gc
}

// init parse cli argumets
func (g *command) init(args []string) error {
	return g.fs.Parse(args)
}

// run executes the command entered
func (g *command) run(client *clientHandler) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	switch g.name() {
	case string(add):
		input := &pb.ToDoItem{
			Title:       *g.flagValues[title].(*string),
			Description: *g.flagValues[description].(*string),
			Status:      pb.Status(*g.flagValues[status].(*int)),
		}
		res, err = client.AddToDo(ctx, input)

	case string(update):
		input := &pb.UpdateToDoItem{
			Hash:   *g.flagValues[hash].(*string),
			Status: pb.Status(*g.flagValues[status].(*int)),
		}
		res, err = client.UpdateToDo(ctx, input)

	case string(deleteOne):
		input := &pb.DeleteToDoItem{
			Hash: *g.flagValues[hash].(*string),
		}
		res, err = client.DeleteToDo(ctx, input)

	case string(getAll):
		res, err = client.GetAllToDo(ctx, &empty.Empty{})
		stream := res.(pb.TodoListService_GetAllToDoClient)
		toDos := make([]*pb.GetToDoItem,0)
		for {
			item,errRecv := stream.Recv()
			if errRecv == io.EOF {
				break
			}
			if errRecv != nil {
				err = errRecv
			}
			toDos = append(toDos,item)
		}
		res = toDos
	default:
		return res, errors.New("command not defined")
	}
	return res, err
}

// name returnes the name of command
func (g *command) name() string {
	return g.fs.Name()
}

// startCli handles the cli
func startCli(client *clientHandler, args []string) (interface{}, error) {
	if len(args) < 1 {
		return nil, errors.New("a sub-command expected")
	}
	cmds := []command{
		newCommand(add, addCmdFlags),
		newCommand(update, updateCmdFlags),
		newCommand(deleteOne, deleteCmdFlags),
		newCommand(getAll, getAllCmdFlags),
		newCommand(help, nil),
	}
	cmdName := os.Args[1]
	if cmdName == string(help) {
		log.Printf("%v", cliUsage)
		return nil, nil
	}
	for _, cmd := range cmds {
		if cmd.name() == cmdName {
			if err := cmd.init(os.Args[2:]); err != nil {
				log.Printf("fail to parse: %v", err)
			}
			return cmd.run(client)
		}
	}
	return nil, fmt.Errorf("unknown cmdName: %s", cmdName)
}

// getAndRunClient create grpc clientHandler
func getAndRunClient(host string) *clientHandler {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	log.Printf("Dialled OK...")
	client := pb.NewTodoListServiceClient(conn)
	return &clientHandler{
		cancel,
		conn.Close,
		client,
	}
}

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Panicf("cannot parse config file: %v", err)
	}
	clientHandler := getAndRunClient(conf.Server.Grpc.Host)
	defer func() {
		clientHandler.cancel()
		if err := clientHandler.close(); err != nil {
			log.Panicf("connection should be closed but got error: %v", err)
		}
	}()
	res, err := startCli(clientHandler, os.Args[1:])
	log.Printf("res: %+v\n, error: %+v\n", res, err)
}
