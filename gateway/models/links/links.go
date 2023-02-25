package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	LinkID          primitive.ObjectID `bson: "_id"`
	LinkDescription string             `bson: "linkDescription" json: "linkDescription" `
	Favorited       bool               `bson: "favorited" json: "favorited"`
	UserIDS         *models.User       `bson: "userIDS" json: "userIDS"`
	OrgID           primitive.ObjectID `bson: "_id, omitempty"`
	PlannerIDS      *models.Planner    `bson: "plannerIDS" json: "plannerIDS"`
}
