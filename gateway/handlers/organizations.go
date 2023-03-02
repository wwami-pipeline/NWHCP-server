package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ==========================================================ORGANIZATION MODELS============================================================
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
	UsersFavorited    []*User            `bson: "usersFavorited" json: "usersFavorited"`
	UsersCompleted    []*User            `bson: "usersCompleted" json: "usersCompleted"`
	UsersCompleting   []*User            `bson: "usersCompleting" json: "usersCompleting"`
	OrgDescription    string             `bson: "orgDescription" json: "orgDescription"`
	StudentsContacted []*User            `bson: "studentsContacted" json: "studentsContacted"`
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

// ===============================================ORGANIZATION CONTROLLERS=========================================================

// assign type to all funcs for easy access in main
// User controller should have access to a Mongo session
type OrganizationController struct {
	session *mongo.Client
}

// convenience function returns UserController
// pass around the address, not the big data structure
func NewOrganizationController(s *mongo.Client) *OrganizationController {
	return &OrganizationController{s}
}

func (oc OrganizationController) CreateOrganization(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	o := &NewOrganization{
		OrgTitle:      params["orgTitle"],
		OrgWebsite:    params["orgWebsite"],
		StreetAddress: params["streetAddress"],
		City:          params["city"],
		State:         params["state"],
		ZipCode:       params["zipCode"],
		Phone:         params["phone"],
		Email:         params["email"],
	}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&o)

	// Create BSON ID
	o.OrgId = primitive.NewObjectID()

	// does this create a collection if it doesn't exist?
	oc.session.Database("mongodb").Collection("organizations").InsertOne(context.TODO(), &o)

	json.NewEncoder(w).Encode(&o)

}

func (oc OrganizationController) GetOrgByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]

	oid, _ := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	fmt.Println("ObjectIDFromHex ERROR:", err)
	// }

	// empty struct to store org
	result := &Organization{}

	// Fetch org
	err := oc.session.Database("mongodb").Collection("surveys").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		fmt.Println("Organization not found")
		os.Exit(1)
		return
	}

	json.NewEncoder(w).Encode(*result)

}
