package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pipeline-db/handlers"
	"github.com/pipeline-db/stores"

	mgo "gopkg.in/mgo.v2"
)

func main() {

	// Set up for microservice information

	portAddr := os.Getenv("PORTADDR") //nwhealthcareerpath.uw.edu:4000 //capstone-test.andrewkan.me:4000
	print(portAddr)
	if len(portAddr) == 0 {
		portAddr = "localhost:4002"
	}
	log.Printf("PORTADDR: %s", portAddr)

	dbAddr := os.Getenv("DBADDR") //pipelineDB:27017
	if len(dbAddr) == 0 {
		dbAddr = "localhost:27017"
	}
	log.Printf("DBADDR: %s", dbAddr)

	mongoSession, err := mgo.Dial(dbAddr)
	if err != nil {
		fmt.Println("Error dialing dbaddr")
	} else {
		fmt.Println("Success!")
	}
	schoolStore, err := stores.NewSchoolStore(mongoSession, "school", "mongodb")
	orgStore, err := stores.NewOrganizationStore(mongoSession, "organization", "mongodb")

	hctx := &handlers.Context{
		Store1: schoolStore,
		Store2: orgStore,
	}

	// jsonfile, _ := ioutil.ReadFile("./data/orgs.json")
	// var orgs []models.Organization
	// err = json.Unmarshal(jsonfile, &orgs)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, org := range orgs {
	// 	dbOrg := &models.Organization{}
	// 	dbOrg.OrgTitle = org.OrgTitle
	// 	dbOrg.OrgWebsite = org.OrgWebsite
	// 	dbOrg.StreetAddress = org.StreetAddress
	// 	dbOrg.City = org.City

	// 	dbOrg.State = org.State
	// 	dbOrg.ZipCode = org.ZipCode
	// 	dbOrg.Phone = org.Phone
	// 	dbOrg.Email = org.Email
	// 	dbOrg.ActivityDesc = org.ActivityDesc
	// 	dbOrg.Lat = org.Lat
	// 	dbOrg.Long = org.Long
	// 	dbOrg.HasShadow = org.HasShadow
	// 	dbOrg.HasCost = org.HasCost
	// 	dbOrg.HasTransport = org.HasTransport
	// 	dbOrg.Under18 = org.Under18
	// 	dbOrg.CareerEmp = org.CareerEmp
	// 	dbOrg.GradeLevels = org.GradeLevels

	// 	orgStore.InsertOrg(dbOrg)
	// 	print(dbOrg)
	// }

	// [{"OrgID": 1, "OrgTitle": "Pre-Health Professions Advising Program",
	// "OrgWebsite": "www.uidaho.edu/pre-health", "StreetAddress": "875 Perimeter Drive",
	// "City": "Moscow", "State": "Idaho", "ZipCode": "83844-2436", "Phone": "(208) 885-5809",
	//  "Email": "pre-health@uidaho.edu",
	//  "ActivityDesc": "The Pre-Health Advising Program at the University of Idaho serves as a resource
	//  for students and alumni, from all majors, who are exploring graduate and professional programs in healthcare.
	//    Services provided include assistance with career exploration, prerequisite course sequencing, advice for
	//     resume building and entrance exam preparation, and support with the application and interview process. "
	// 	, "Lat": 46.727471200000004, "Long": -117.02390190000001,
	// 	"HasShadow": false, "HasCost": false, "HasTransport": false, "Under18": false,
	// 	"CareerEmp": ["Medicine", "Dentistry", "Pharmacy", "Public Health", "Generic Health Sciences",
	// 	"Allied Health", "STEM"],
	// 	"GradeLevels": [9, 10, 11, 12]},

	/* for schools trying to get other schools working rn
	jsonfile, _ := ioutil.ReadFile("./data/wa-schools-clean.json")
	var schools []models.School
	err = json.Unmarshal(jsonfile, &schools)
	if err != nil {
		panic(err)
	}
	for _, school := range schools {
		dbSchool := &models.School{}
		dbSchool.SchoolName = school.SchoolName
		dbSchool.Street = school.Street
		dbSchool.City = school.City
		dbSchool.County = school.County
		// Issue with FullAddress not going in, taking out for now. Address can be built with
		// above components
		// dbSchool.FullAddress = school.FullAddress
		dbSchool.Lat = school.Lat
		dbSchool.Lng = school.Lng
		dbSchool.SchoolDistrictName = school.SchoolDistrictName
		dbSchool.Zip = school.Zip
		schoolStore.InsertSchool(dbSchool)
		print(dbSchool)
	}

	*/
	//ignore compiler
	log.Printf("%v", hctx)
	log.Printf("%v", schoolStore)
	log.Printf("%v", orgStore)

	// TLSCERT := os.Getenv("TLSCERT")
	// TLSKEY := os.Getenv("TLSKEY")

	if err != nil {
		fmt.Println("Error creating store")
	}

	mux := http.NewServeMux()
	fmt.Println("Pipeline-DB Microservice")
	mux.HandleFunc("/v1/pipeline-db/populate-school", hctx.PopulateDB)
	mux.HandleFunc("/v1/pipeline-db/populate-organizations", hctx.PopulateDB2)

	mux.HandleFunc("/", hctx.TestHandler)

	mux.HandleFunc("/v1/pipeline-db/getallschools", hctx.SchoolHandler)
	mux.HandleFunc("/v1/pipeline-db/getallorgs", hctx.OrgHandler)

	log.Printf("server listening at http://%s...", portAddr)
	// log.Fatal(http.ListenAndServeTLS(portAddr, TLSCERT, TLSKEY, mux))
	log.Fatal(http.ListenAndServe(portAddr, mux))
}
