package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type EStatus int

const (
	TODO       EStatus = 0
	INPROGRESS EStatus = 1
	DONE       EStatus = 2
)

type ToDo struct {
	Id          primitive.ObjectID `bson:"_id,omitepty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Status      EStatus            `bson:"status"`
}
