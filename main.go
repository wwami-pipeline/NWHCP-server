package main

import (
	"encoding/json"
	"fmt"
	"log"
	"pipeline-db/models"
	"pipeline-db/stores"

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

	//// Script to insert schools
	// jsonFile, err := ioutil.ReadFile("/Users/studentuser/Desktop/wa-schools-clean.json")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Successfully Opened wa-schools.json")
	// // defer the closing of our jsonFile so that we can parse it later on

	// var schools []models.School
	// json.Unmarshal(jsonFile, &schools)

	// fmt.Println(schools[0])

	// for _, school := range schools {
	// 	dbSchool := &models.School{}
	// 	dbSchool.SchoolName = school.SchoolName
	// 	dbSchool.Street = school.Street
	// 	dbSchool.City = school.City
	// 	dbSchool.County = school.County
	// 	dbSchool.FullAddress = school.FullAddress
	// 	dbSchool.Latitude = school.Latitude
	// 	dbSchool.Longitude = school.Longitude
	// 	dbSchool.SchoolDistrictName = school.SchoolDistrictName
	// 	dbSchool.Zip = school.Zip
	// 	schoolStore.InsertSchool(dbSchool)
	// }

	// test, _ := schoolStore.GetBySchoolName("Federal Way High School")
	// fmt.Println(test)

	// Can add a json object for school.
	school := &models.School{}
	s := `{"school_name":"Beezley Springs Elementary","school_district_name":"Ephrata School District","full_address":"501 C ST NW  EPHRATA Washington 98823-0000","street":"501 C ST NW","city":"EPHRATA","county":"Grant County","state":"Washington","zip":"98823-0000","lat":"47.3259809","lng":"-119.5496231"}`
	buffer := []byte(s)

	json.Unmarshal(buffer, school)

	fmt.Println(school.SchoolName)

	fmt.Println(school.Latitude)

	_, _ = schoolStore.InsertSchool(school)

	// fmt.Printf("inserted: %v", inserted)

	// mux := http.NewServeMux()
	fmt.Println("SchoolDatabase Microservice")

	// log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServe(portAddr, mux))
}
