package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Planner struct {
	PlannerID          primitive.ObjectID     `bson: "_id"`
	IsMonthlyPlanner   bool                   `bson: "isMonthlyPlanner" json: "isMonthlyPlanner"`
	IsYearlyPlanner    bool                   `bson: "isYearlyPlanner" json: "isYearlyPlanner"`
	IsAcademicPlanner  bool                   `bson: "isAcademicPlanner" "isAcademicPlanner"`
	NotesIDS           []*models.Note         `bson: "notesIDS" json: "notesIDS"`
	OrgIDS             []*models.Organization `bson: "orgIDS" json: "orgIDS"`
	UserID             primitive.ObjectID     `bson: "_id"`
	IsFallPlanner      bool                   `bson: "isFallPlanner" json: "isFallPlanner"`
	IsWinterPlanner    bool                   `bson: "isWinterPlanner" json: "isWinterPlanner"`
	IsSpringPlanner    bool                   `bson: "isSpringPlanner" json: "isSpringPlanner"`
	IsSummerPlanner    bool                   `bson: "isSummerPlanner" json: "isSummerPlanner"`
	LinkIDS            []*models.Link         `bson: "linkIDS" json: "linkIDS"`
	DateCreated        string                 `bson: "dateCreated" json: "dateCreated"`
	PlannerDescription string                 `bson: "plannerDescription" json: "plannerDescription"`
}
