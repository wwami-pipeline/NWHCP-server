package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

// ====================================================================USER MODELS========================================================

// bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

// User represents a user account in the database (student)
// how do I store the stuff?
type User struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Email     string             `bson:"email" json:"email"`       //never JSON encoded/decoded
	PassHash  []byte             `bson:"password" json:"password"` //never JSON encoded/decoded
	UserName  string             `bson: "userName" json:"userName"`
	FirstName string             `bson: "firstName" json: "firstName"`
	JoinDate  string             `bson: "joinDate" json:"joinDate"`
	State     string             `bson: "state" json:"state"`
	// AllOrgs   *UserOrgs          `bson: "allOrgs" json: "allOrgs"`
	FavoritedOrganizations []*OrgID     `bson: "favoritedOrganizations" json:"favoritedOrganizations"`
	CompletedPrograms      []*OrgID     `bson: "completedPrograms" json:"completedPrograms"`
	InProcessPrograms      []*OrgID     `bson: "inProcessPrograms" json:"inProcessPrograms"`
	PathwayPrograms        []*OrgID     `bson: "pathwayPrograms" json: "pathwayPrograms"`
	AcademicPrograms       []*OrgID     `bson: "academicPrograms" json: "academicPrograms"`
	Notes                  []*NoteID    `bson: "notes" json:"notes"`
	Links                  []*LinkID    `bson: "links" json: "links"`
	Planners               []*PlannerID `bson: "planners" json: "planners"`
	QuantityPlanners       int          `bson: "quantityPlanners" json: "quantityPlanners"`
	OrgsContacted          []*OrgID     `bson: "orgsContacted" json: "orgsContacted"`
}

type UserOrgs struct {
	FavoritedOrganizations        []*OrgID `bson: "favoritedOrganizations" json: "favoritedOrganizations"`
	CompletedOrganizations        []*OrgID `bson: "completedOrganizations" json:"completedOrganizations"`
	InProgressOrganizations       []*OrgID `bson: "inProgressOrganizations" json:"inProgressOrganizations"`
	PathwayOrganizations          []*OrgID `bson: "pathwayOrganizations" json: "pathwayOrganizations"`
	QuantityPathwayOrganizations  int      `bson: "quantityPathwayOrganizations" json: "quantityPathwayOrganizations"`
	AcademicOrganizations         []*OrgID `bson: "academicOrganizations" json: "academicOrganizations"`
	QuantityAcademicOrganizations int      `bson: "quantityAcademicOrganizations" json: "quantityAcademicOrganizations"`
}

// used in in many-to-many relationship modeling
type UserID struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
}

// Credentials represents user sign-in credentials (student)
type Credentials struct {
	Email    string `bson: "email" json:"email"`
	Password string `bson: "password" json:"password"`
}

// NewUser represents a new user signing up for an account (student)
type NewUser struct {
	ID           primitive.ObjectID `"bson:_id"`
	Email        string             `bson: "email" json:"email"`
	PassHash     []byte             `bson:"password" json:"password"` //never JSON encoded/decoded
	PasswordConf []byte             `bson: "passwordConf" json:"passwordConf"`
	FirstName    string             `bson: "firstName" json:"firstName"`
	UserName     string             `bson: "userName" json:"userName"`
}

// Updates represents updates allowed to be edited by user (student)
type Updates struct {
	FirstName      string      `bson: "firstName" json:"firstName"`
	UserName       string      `bson: "userName" json:"userName"`
	UpdateUserOrgs []*UserOrgs `bson: "updateUserOrgs" json: "updateUserOrgs"`
	// FavoritedOrganizations []*Organization `bson: "favoritedOrganizations" json:"favoritedOrganizations"`
	// CompletedPrograms      []*Organization `bson: "completedPrograms" json:"completedPrograms"`
	// InProcessPrograms      []*Organization `bson: "inProcessPrograms" json:"inProcessPrograms"`
	// UserPathwayPrograms    []*Organization `bson: "userPathwayPrograms" json: "userPathwayPrograms" `
	// UserAcademicPrograms   []*Organization `bson: "userAcademicPrograms" json: "userAcademicPrograms"`
	UserNotes    []*NoteID    `bson: "userNotes" json:"userNotes"`
	UserLinks    []*LinkID    `bson: "userLinks" json: "userLinks"`
	UserPlanners []*PlannerID `bson: "userPlanners" json: "userPlanners"`
}

// Notes represents users' (student) notes about programs
type Note struct {
	UserID          primitive.ObjectID `bson:"_id, omitempty"`
	NoteID          primitive.ObjectID `bson: "_id"`
	NoteContent     string             `bson: "noteContent" json:noteContent`
	OrgID           primitive.ObjectID `bson: "_id"`
	CreatedAt       string             `bson: "createdAt" json:createdAt`
	UpdatedAt       string             `bson: "updatedAt" json:updatedAt`
	Reviewed        bool               `bson: "reviewed" json:reviewed`
	NoteDescription string             `bson: "noteDescription" json: "noteDescription"`
}

type NoteID struct {
	NoteID primitive.ObjectID `bson: "noteID" json: "noteID"`
}

type Link struct {
	LinkID          primitive.ObjectID `bson: "_id"`
	LinkDescription string             `bson: "linkDescription" json: "linkDescription" `
	Favorited       []*UserID          `bson: "favorited" json: "favorited"`
	UserIDs         []*UserID          `bson:"userIDS" json: "userIDS"`
	PlannerIDs      []*PlannerID       `bson: "plannerIDS json: "plannerIDS"`
	OrgID           primitive.ObjectID `bson: "_id, omitempty"`
	NoteIDs         []*NoteID          `bson: "noteIDS" json: "noteIDS"`
}

type LinkID struct {
	LinkID primitive.ObjectID `bson: "linkID" json: "linkID"`
}

type Planner struct {
	PlannerID          primitive.ObjectID `bson: "_id"`
	IsMonthlyPlanner   bool               `bson: "isMonthlyPlanner" json: "isMonthlyPlanner"`
	IsYearlyPlanner    bool               `bson: "isYearlyPlanner" json: "isYearlyPlanner"`
	IsAcademicPlanner  bool               `bson: "isAcademicPlanner" "isAcademicPlanner"`
	NotesIDS           []*NoteID          `bson: "notesIDS" json: "notesIDS"`
	OrgIDS             []*OrgID           `bson: "orgIDS" json: "orgIDS"`
	UserID             primitive.ObjectID `bson: "_id"`
	IsFallPlanner      bool               `bson: "isFallPlanner" json: "isFallPlanner"`
	IsWinterPlanner    bool               `bson: "isWinterPlanner" json: "isWinterPlanner"`
	IsSpringPlanner    bool               `bson: "isSpringPlanner" json: "isSpringPlanner"`
	IsSummerPlanner    bool               `bson: "isSummerPlanner" json: "isSummerPlanner"`
	LinkIDS            []*LinkID          `bson: "linkIDS" json: "linkIDS"`
	DateCreated        string             `bson: "dateCreated" json: "dateCreated"`
	PlannerDescription string             `bson: "plannerDescription" json: "plannerDescription"`
}

type PlannerID struct {
	PlannerID primitive.ObjectID `bson: "plannerID" json: "plannerID"`
}

// ==================================================USER CONTROLLERS=================================================================

// assign type to all funcs for easy access in main
// User controller should have access to a Mongo session
type UserController struct {
	session *mongo.Client
}

// convenience function returns UserController
// pass around the address, not the big data structure
func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

// refactor this to read user input from the body... - to do
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	pw := params["password"]
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pw), 14)

	err := bcrypt.CompareHashAndPassword([]byte(bytes), []byte(pw))
	if err != nil {
		fmt.Println("Password hashing didn't seem to work...")
	}

	// check that password and password confirmation match - to do
	// pwc := params["passwordConf"]
	// pwcBytes, _ := bcrypt.GenerateFromPassword([]byte(pwc), 14)

	// error if passwords don't match

	// if bytes != pwcBytes {
	// 	fmt.Println("Passwords don't match")
	// 	return
	// }

	u := &NewUser{
		Email:        params["email"],
		PassHash:     bytes,
		PasswordConf: bytes,
		FirstName:    params["firstName"],
		UserName:     params["userName"],
	}
	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// Create BSON ID
	u.ID = primitive.NewObjectID()

	uc.session.Database("mongodb").Collection("usersTest").InsertOne(context.TODO(), &u)

	json.NewEncoder(w).Encode(&u)

}

func (uc UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("primitive.ObjectIDFromHex ERROR:", err)
	}

	res, err := uc.session.Database("mongodb").Collection("usersTest").DeleteOne(context.TODO(), bson.M{"_id": oid})
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

func (uc UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	oid, _ := primitive.ObjectIDFromHex(id)

	result := &User{}

	err := uc.session.Database("mongodb").Collection("usersTest").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		fmt.Println("User not found")
		os.Exit(1)
		return
	}

	json.NewEncoder(w).Encode(*result)
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := uc.session.Database("mongodb").Collection("usersTest").Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all Users ERROR:", err)
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

/* =================================================UPDATE USER ORGS====================================================*/

// add orgs to favorite by reference
func (uc UserController) AddOrgToFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgsid"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	u_selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	u_change := bson.M{
		"$addToSet": bson.M{
			"favoritedOrganizations": ooid,
		},
	}

	// update user doc
	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), u_selector, u_change)

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": ooid,
		},
	}

	o_change := bson.M{
		"$addToSet": bson.M{
			"favorited": oid,
		},
	}
	// update org doc
	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)

}

func (uc UserController) DeleteOrgFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgsid"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	change := bson.M{
		"$pull": bson.M{
			"favoritedOrganizations": ooid,
		},
	}

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": ooid,
		},
	}

	o_change := bson.M{
		"$pull": bson.M{
			"favorited": oid,
		},
	}

	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)

	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)
}

func (uc UserController) AddToPathwayOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")
	// no body - JSON can't accept ObjectIDs
	w.Header().Set("Content-length", "0")

	params := mux.Vars(r)

	id := params["id"]
	orgId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add org reference id to pathwayPrograms
	change := bson.M{
		"$addToSet": bson.M{
			"pathwayPrograms": ooid,
		},
	}

	// update user doc
	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)

}
func (uc UserController) AddToAcademicOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")
	// no body - JSON can't accept ObjectIDs
	w.Header().Set("Content-length", "0")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	change := bson.M{
		"$addToSet": bson.M{
			"academicPrograms": ooid,
		},
	}

	// update user doc
	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)
}
func (uc UserController) AddToCompletedOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")
	// no body - JSON can't accept ObjectIDs
	w.Header().Set("Content-length", "0")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	u_selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	u_change := bson.M{
		"$addToSet": bson.M{
			"completedPrograms": ooid,
		},
	}

	// update user doc
	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), u_selector, u_change)

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": ooid,
		},
	}

	o_change := bson.M{
		"$addToSet": bson.M{
			"completed": oid,
		},
	}
	// update org doc
	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)
}
func (uc UserController) AddToinProgressOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")
	// no body - JSON can't accept ObjectIDs
	w.Header().Set("Content-length", "0")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	u_selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	u_change := bson.M{
		"$addToSet": bson.M{
			"inProgressPrograms": ooid,
		},
	}

	// update user doc
	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), u_selector, u_change)

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": ooid,
		},
	}

	o_change := bson.M{
		"$addToSet": bson.M{
			"inProgress": oid,
		},
	}
	// update org doc
	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)
}

func (uc UserController) DeleteFromPathwayOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	change := bson.M{
		"$pull": bson.M{
			"pathwayPrograms": ooid,
		},
	}

	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)

}
func (uc UserController) DeleteFromAcademicOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	change := bson.M{
		"$pull": bson.M{
			"academicPrograms": ooid,
		},
	}

	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)

}
func (uc UserController) DeleteFromCompletedOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	change := bson.M{
		"$pull": bson.M{
			"completedPrograms": ooid,
		},
	}

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": ooid,
		},
	}

	o_change := bson.M{
		"$pull": bson.M{
			"completed": oid,
		},
	}

	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)

	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)
}
func (uc UserController) DeleteFrominProgressOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]
	orgsId := params["orgId"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	ooid, err := primitive.ObjectIDFromHex(orgsId)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add orgId to favorites
	change := bson.M{
		"$pull": bson.M{
			"inProgressPrograms": ooid,
		},
	}

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": ooid,
		},
	}

	o_change := bson.M{
		"$pull": bson.M{
			"inProgress": oid,
		},
	}

	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), selector, change)

	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)
}

//====================================================UPDATE USER NOTES===========================================================

// proper POST request
// this function to be refactored; new name "GetNoteBody"
func (uc UserController) CreateNote(r *http.Request, n *Note) error {

	if r.Body == nil {
		return errors.New("request body is empty")
	}

	if n == nil {
		return errors.New("a note is required")
	}

	postBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// unmarshall data from request body - this to be returned and function to end here after refactor
	json.Unmarshal(postBody, &n)

	// this goes in separate post function; new name "NoteToPostNote"

	// Create BSON ID
	n.NoteID = primitive.NewObjectID()

	// create new note - this should be in separate function
	uc.session.Database("mongodb").Collection("notes").InsertOne(context.TODO(), &n)

	// encode/decode for sending/receiving JSON to/from a stream
	return json.NewDecoder(r.Body).Decode(&n)

}

func (uc UserController) AddNoteID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	// Resource in URI
	// userId
	// orgId
	// noteId

	params := mux.Vars(r)

	u := params["userId"]
	o := params["orgId"]
	n := params["noteId"]

	uid, err := primitive.ObjectIDFromHex(u)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	oid, err := primitive.ObjectIDFromHex(o)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	nid, err := primitive.ObjectIDFromHex(n)
	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR:", err)
	}

	// grab user
	u_selector := bson.M{
		"_id": bson.M{
			"$eq": uid,
		},
	}

	// add note reference to user notes
	u_change := bson.M{
		"$addToSet": bson.M{
			"notes": nid,
		},
	}

	// grab org
	o_selector := bson.M{
		"_id": bson.M{
			"$eq": oid,
		},
	}

	// add note reference to org notes
	o_change := bson.M{
		"$addToSet": bson.M{
			"orgNotes": nid,
		},
	}

	// grab note
	n_selector := bson.M{
		"_id": bson.M{
			"$eq": nid,
		},
	}

	// update note with orgID && userID
	n_change := bson.M{
		"$addToSet": bson.M{
			"userID": uid,
			"orgID":  oid,
		},
	}

	// associate note with user
	uc.session.Database("mongodb").Collection("usersTest").UpdateOne(context.TODO(), u_selector, u_change)

	// associate note with organization
	uc.session.Database("mongodb").Collection("organizations").UpdateOne(context.TODO(), o_selector, o_change)

	// update note
	uc.session.Database("mongodb").Collection("notes").UpdateOne(context.TODO(), n_selector, n_change)

}
