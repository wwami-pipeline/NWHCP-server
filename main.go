package main

import (
	"fmt"
	"log"
	"pipeline-db/models"
	"pipeline-db/stores"

	"gopkg.in/mgo.v2/bson"

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
		SchoolName: "Test3",
		SchoolID:   testID,
	}
	// asdf := bson.ObjectIdHex("5acef803aa50a545aa77ff7a")
	log.Printf("TestID: %v", testID)

	insertSchool, err := schoolStore.InsertSchool(testSchool)
	log.Printf("insertSchool: %v", insertSchool)

	getSchool, err := schoolStore.GetByID(testID)
	log.Printf("GetSchoolByID: %v", getSchool)

	getSchoolByName, err := schoolStore.GetBySchoolName("Test3")
	log.Printf("GetSchoolByName before remove: %v", getSchoolByName)

	// removeSchool := schoolStore.DeleteSchool(testID)
	// log.Printf("Remove : %v", removeSchool)

	updateSchool := &models.UpdatedSchool{
		SchoolName: "Test4",
	}

	updated := schoolStore.UpdateSchool(testSchool.SchoolName, updateSchool)
	log.Printf("Updated: %v", updated)

	// getSchoolByNameAfterRemove, err := schoolStore.GetBySchoolName("Test3")
	// log.Printf("GetSchoolByName after remove: %v", getSchoolByNameAfterRemove)

	// log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServe(portAddr, mux))
}
