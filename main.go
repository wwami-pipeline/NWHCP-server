package main

import (
	"fmt"
	"log"

	"github.com/pipeline-db/models"
	"gopkg.in/mgo.v2/bson"

	"github.com/pipeline-db/stores"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	// dbAddr := os.Getenv("DBADDR")
	// if len(dbAddr) == 0 {
	// 	dbAddr = "school_mongo:27017"
	// }
	// log.Printf("DBADDR: %s", dbAddr)

	mongoSession, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Error dialing dbaddr")
	} else {
		fmt.Println("Success!")
	}

	schoolStore, err := stores.NewSchoolStore(mongoSession, "school", "mongodb")

	//ignore compiler
	log.Printf("%v", schoolStore)

	if err != nil {
		fmt.Println("Error creating store")
	}

	// mux := http.NewServeMux()
	fmt.Println("SchoolDatabase Microservice")

	testID := bson.NewObjectId()
	testSchool := &models.School{
		SchoolName: "Test",
		SchoolID:   testID,
	}

	insertSchool, err := schoolStore.InsertSchool(testSchool)
	log.Printf("insertSchool: %v", insertSchool)

	getSchool, err := schoolStore.GetByID(testID)
	log.Printf("GetSchool: %v", getSchool)

	// log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServe(portAddr, mux))
}
