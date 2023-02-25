package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nwhcp/nwhcp-server/gateway/models"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/mongo"
)

// return one, rest optional
type ClientController struct {
	oc OrganizationController
	uc UserController
	pc PlannerController
	lc LinkController
}

// controllers called from main exe file
type OrganizationController struct {
	session *mongo.Client
}
type UserController struct {
	session *mongo.Client
}

type PlannerController struct {
	session *mongo.Client
}

type LinkController struct {
	session *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func NewOrganizationController(s *mongo.Client) *OrganizationController {
	return &OrganizationController{s}
}
func NewPlannerController(s *mongo.Client) *PlannerController {
	return &PlannerController{s}
}

func NewLinkController(s *mongo.Client) *LinkController {
	return &LinkController{s}
}

// return one, rest optional
type ClientType struct {
	um  *models.User
	om  *models.Organization
	pm  *models.Planner
	lm  *models.Link
	mnu *models.NewUser
}

type RequestModel interface {
	ReqRes(cc *ClientController)
}

type Client interface {
	// receives a type, client controller
	createNew() RequestModel // diff logic for each controller
	deleteOne() RequestModel
	deleteMany() RequestModel
	updateOne() RequestModel
	updateMany() RequestModel
	getOneByID() RequestModel
	getAllByID() RequestModel
}

/* User Controllers
- createNew() - creates a new user
- deleteOne() - deletes a user
- updateOne() - updates a field of a user
- updateMany() - updates many fields of a user
- getOneByID() - get a user by ID
- getAllByID() - get all users by ID
*/

func (uc UserController) CreateNew(w http.ResponseWriter, r *http.Request) *ClientType {

	params := mux.Vars(r)

	pw := params["password"]
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pw), 14)

	err := bcrypt.CompareHashAndPassword([]byte(bytes), []byte(pw))
	if err != nil {
		fmt.Println("Password hashing didn't seem to work...")
	}

	u := &models.NewUser{
		Email:        params["email"],
		PassHash:     bytes,
		PasswordConf: bytes,
		FirstName:    params["firstName"],
		UserName:     params["userName"],
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	json.NewDecoder(r.Body).Decode(&u)
	return &u
}

/* Organization Controllers
- createNew() - creates a new organization
- deleteOne() - deletes a organization
- updateOne() - updates a field of a organization
- updateMany() - updates many fields of a organization
- getOneByID() - get a organization by ID
- getAllByID() - get all organizations by ID
*/

/* Planner Controllers
- createNew() - creates a new planner
- deleteOne() - deletes a planner
- updateOne() - updates a field of a planner
- updateMany() - updates many fields of a planner
- getOneByID() - get a planner by ID
- getAllByID() - get all planners by ID
*/

/* Link Controllers
- createNew() - creates a new link
- deleteOne() - deletes a link
- updateOne() - updates a field of a link
- updateMany() - updates many fields of a link
- getOneByID() - get a link by ID
- getAllByID() - get all links by ID
*/
