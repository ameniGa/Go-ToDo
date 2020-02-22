package testData

import (
	entity "github.com/3almadmoon/ameni-assignment/entities"
	"github.com/google/uuid"
)

// TTCreateHandler represents the table test for create handler
var TTCreateHandler = []struct {
	Name     string
	Type     string
	Uri string
	HasError bool
}{
	{"supported type", "mongo","mongodb://localhost:27017", false},
	{"unsupported type", "notType","" ,true},
	{"unsupported uri", "mongo","fakeURI", true},
}

var Thash = uuid.New().String()

// TTtoDo represents the table test for to do CRUD
var TTtoDo = []struct{
	Name             string
	Hash             string
	Title            string
	Description      string
	Status           entity.EStatus
	IsValidCtx       bool
	HasErrorOnCreate bool
	HasErrorOnUpdate bool
}{
	{"empty Title",Thash,"","learn interfaces",entity.TODO,true,true,false},
	{"empty Hash","","","",entity.TODO,true,true,true},
	{"invalid context","","","",entity.TODO,false,true,true},
	{"proper item",Thash,"learn go","learn interfaces",entity.TODO,true,false,false},
	{"existing item",Thash,"learn go","learn interfaces",entity.INPROGRESS,true,true,false},
}
