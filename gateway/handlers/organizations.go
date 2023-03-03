package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

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
	// UsersFavorited    []*User            `bson: "usersFavorited" json: "usersFavorited"`
	// UsersCompleted    []*User            `bson: "usersCompleted" json: "usersCompleted"`
	// UsersCompleting   []*User            `bson: "usersCompleting" json: "usersCompleting"`
	AllUsers          *OrgUsers    `bson: "allUsers" json: "allUsers"`
	OrgDescription    string       `bson: "orgDescription" json: "orgDescription"`
	StudentsContacted []*UserID    `bson: "studentsContacted" json: "studentsContacted"`
	Tags              []string     `bson: "tags" json: "tags"`
	OrgPlanners       []*PlannerID `bson: "orgPlanners" json: "orgPlanners"`
	OrgLinks          []*LinkID    `bson: "orgLinks" json: "orgLinks"`
	OrgNotes          []*NoteID    `bson: "orgNotes" json: "orgNotes"`
}

type OrgID struct {
	OrgId primitive.ObjectID `bson:"_id"`
}

type OrgUsers struct {
	Favorited  []*UserID `bson: "favorited" json: "favorited"`
	InProgress []*UserID `bson: "inProgress" json: "inProgress"`
	Completed  []*UserID `bson: "completed" json: "completed"`
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

	params := mux.Vars(r)

	id := params["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	oid, _ := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	fmt.Println("ObjectIDFromHex ERROR:", err)
	// }

	// empty struct to store org
	result := &Organization{}

	// Fetch org
	err := oc.session.Database("mongodb").Collection("organizations").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		fmt.Println("Organization not found")
		os.Exit(1)
		return
	}

	json.NewEncoder(w).Encode(*result)

}

func (oc OrganizationController) GetOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := oc.session.Database("mongodb").Collection("organizations").Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all Organizations ERROR:", err)
	}

	// collect results
	for cursor.Next(context.TODO()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println("Cursor.Next() ERROR:", err)
			os.Exit(1)
		}
		json.NewEncoder(w).Encode(result)
		fmt.Println("\nresult type:", reflect.TypeOf(result))
		fmt.Println("result:", result)
	}

}

func (oc OrganizationController) DeleteOrganizationByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	res, err := oc.session.Database("mongodb").Collection("organizations").DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		log.Fatal("DeleteOne() ERROR:", err)
	}
	if res.DeletedCount == 0 {
		fmt.Println("Delete One() document not found:", res)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Print results of DeleteOne method
	fmt.Println("DeleteOne Result:", res)
	json.NewEncoder(w).Encode(res)

}
