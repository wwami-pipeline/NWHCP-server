package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	UserID          primitive.ObjectID `bson:"_id, omitempty"`
	NoteID          primitive.ObjectID `bson: "_id"`
	NoteContent     string             `bson: "noteContent" json:noteContent`
	CreatedAt       string             `bson: "createdAt" json:createdAt`
	UpdatedAt       string             `bson: "updatedAt" json:updatedAt`
	Reviewed        bool               `bson: "reviewed" json:reviewed`
	NoteDescription string             `bson: "noteDescription" json: "noteDescription"`
}
