package handlers

// user controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nwhcp/nwhcp-server/gateway/models/users"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

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

	u := &users.NewUser{
		Email:        params["email"],
		PassHash:     bytes,
		PasswordConf: bytes,
		FirstName:    params["firstName"],
		LastName:     params["lastName"],
		Age:          params["age"],
		State:        params["state"],
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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
	fmt.Println("User Deleted:", res)
	fmt.Println("DeleteOne TYPE:", reflect.TypeOf(res))
}

func (uc UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id := params["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	oid, _ := primitive.ObjectIDFromHex(id)

	result := &users.User{}

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
