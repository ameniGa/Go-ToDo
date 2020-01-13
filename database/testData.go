package database

import(
	entity "github.com/3almadmoon/ameni-assignment/entities"
	"github.com/google/uuid"
)

var Thash = uuid.New().String()
var TtToDo = []struct{
	Name             string
	Hash             string
	Title            string
    Description      string
	Status           entity.EStatus
	IsValidCtx       bool
	HasErrorOnCreate bool
	HasErrorOnUpdate bool
}{
	{"empty title",Thash,"","learn interfaces",entity.TODO,true,true,false},
	{"empty hash","","","",entity.TODO,true,true,true},
	{"invalid context","","","",entity.TODO,false,true,true},
	{"proper item",Thash,"learn go","learn interfaces",entity.TODO,true,false,false},
	{"existing item",Thash,"learn go","learn interfaces",entity.INPROGRESS,true,true,false},
}
