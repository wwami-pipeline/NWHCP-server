package models

import (
	"nwhcp/nwhcp-server/gateway/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Organization represents information for a new org
type Organization struct {
	OrgId             primitive.ObjectID `bson:"_id"`
	OrgTitle          string             `bson:"orgTitle" json: "orgTitle"`
	OrgWebsite        string             `bson:"orgWebsite" json: "orgWebsite"`
	StreetAddress     string             `bson:"streetAddress" json: "streetAddress"`
	City              string             `bson:"city" json: "city"`
	State             string             `bson:"state" json: "state"`
	ZipCode           string             `bson:"zipCode" json: "zipCode"`
	Phone             string             `bson:"phone" json: "phone"`
	Email             string             `bson:"email" json: "email"`
	ActivityDesc      string             `bson:"activityDesc" json: "activityDesc"`
	Lat               float64            `bson:"lat" json: "lat"`
	Long              float64            `bson:"long" json: "long"`
	HasShadow         bool               `bson:"hasShadow" json: "hasShadow"`
	HasCost           bool               `bson:"hasCost" json: "hasCost"`
	HasTransport      bool               `bson:"hasTransport" json: "hasTransport"`
	Under18           bool               `bson:"under18" json: "under18"`
	CareerEmp         []string           `bson:"careerEmp" json: "careerEmp"`
	GradeLevels       []int              `bson:"gradeLevels" json: "gradeLevels"`
	IsPathwayProgram  bool               `bson: "isPathwayProgram" json: "isPathwayProgram"`
	IsAcademicProgram bool               `bson: "isAcademicProgram" json: "isAcademicProgram"`
	UsersFavorited    []*models.User     `bson: "usersFavorited" json: "usersFavorited"`
	UsersCompleted    []*models.User     `bson: "usersCompleted" json: "usersCompleted"`
	UsersCompleting   []*models.User     `bson: "usersCompleting" json: "usersCompleting"`
	OrgDescription    string             `bson: "orgDescription" json: "orgDescription"`
	StudentsContacted []*models.User     `bson: "studentsContacted" json: "studentsContacted"`
	Tags              []string           `bson: "tags" json: "tags"`
}

// OrgInfo represents organization information for buiding searching criteria
type OrgInfo struct {
	SearchContent string   `bson: "searchContent" json: "searchContent"`
	HasShadow     bool     `bson: "hasShadow" json: "hasShadow"`
	HasCost       bool     `bson: "hasCost" json: "hasCost"`
	HasTransport  bool     `bson: "hasTransport" json: "hasTransport"`
	Under18       bool     `bson: "under18" json: "under18"`
	CareerEmp     []string `bson: "careerEmp" json: "careerEmp"`
	GradeLevels   []int    `bson: "gradeLevels" json: "gradeLevels"`
}

type NewOrganization struct {
	OrgId         primitive.ObjectID `bson:"_id"`
	OrgTitle      string             `bson:"orgTitle" json: "orgTitle"`
	OrgWebsite    string             `bson:"orgWebsite" json: "orgWebsite"`
	StreetAddress string             `bson:"streetAddress" json: "streetAddress"`
	City          string             `bson:"city" json: "city"`
	State         string             `bson:"state" json: "state"`
	ZipCode       string             `bson:"zipCode" json: "zipCode"`
	Phone         string             `bson:"phone" json: "phone"`
	Email         string             `bson:"email" json: "email"`
}
