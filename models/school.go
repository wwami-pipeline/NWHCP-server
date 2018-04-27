package models

import (
	"gopkg.in/mgo.v2/bson"
)

type AllSchools struct {
	KeyName    string   `json:"all_schools_key"`
	AllSchools []School `json: "schools"`
	// AllSchool string
}

type School struct {
	SchoolID           bson.ObjectId `json:"schoolID" bson:"_id"`
	SchoolName         string        `json:"school_name"`
	SchoolDistrictName string        `json:"school_district_name"`
	Street             string        `json: "street"`
	City               string        `json: "city"`
	County             string        `json: "county"`
	Zip                string        `json: "zip"`
	Lng                string        `json: "lng"`
	Lat                string        `json: "lat"`
	// FullAddress        string        `json: "full_address, string"`

}

type UpdateSchool struct {
	SchoolName         string `json:"school_name"`
	SchoolDistrictName string `json:"school_district_name"`
	Street             string `json: "street"`
	City               string `json: "city"`
	County             string `json: "county"`
	Zip                string `json: "zip"`
	Lng                string `json: "lng"`
	Lat                string `json: "lat"`
	// FullAddress        string `json: "full_Address"`

}

// [{"school_name":"Grant School","district_name":"Grant Elementary",
// 	"full_address":"811 E Orr Dillon MT 59725","street":"811 E Orr",
// 	"city":"Dillon","state":"MT","county":"Beaverhead","zip":"59725",
// 	"lat":45.0077,"lng":-113.0644},

// [{"school_name":"Anderson School","district_name":"Denali Borough School District",
// 	"full_address":"Box 3120 Anderson AK 99744","street":"Box 3120","city":"Anderson",
// 	"state":"AK","county":"Denali Borough","zip":99744,"lat":64.3437,"lng":-149.1908},

// [{"school_name":"Burns Jr & Sr High School","district_name":"Laramie County School District #2",
// 	"full_address":"305 East County Rd. 213 Burns Wyoming 82053","street":"305 East County Rd. 213",
// 	"city":"Burns","state":"Wyoming","county":"County Road 213","zip":82053,"lat":41.1872,
// 	"lng":-104.3508},
// { "_id" : ObjectId("5ada683a2291a91970975d6d"), "schoolname" : "Wilson Creek Elementary",
// "schooldistrictname" : "Wilson Creek School District", "street" : "PO Box 46",
// "city" : "Wilson Creek", "county" : "Grant County", "zip" : "98860-0000",
// "lng" : "-119.1208498", "lat" : "47.4232005" }
/*
{ "_id" : ObjectId("5ada665b2291a917167249d9"), "schoolname" : "Hoquiam High School",
	"schooldistrictname" : "Hoquiam School District", "fulladdress" : "",
	 "street" : "501 W. Emerson", "city" : "Hoquiam", "county" : "Grays Harbor County",
	  "zip" : "98550-0000", "lng" : "-123.910833", "lat" : "46.982679" }
*/

/*[{"school_name":"Beezley Springs Elementary","school_district_name":"Ephrata School District",
"full_address":"501 C ST NW  EPHRATA Washington 98823-0000","street":"501 C ST NW","city":
"EPHRATA","county":"Grant County","state":"Washington","zip":"98823-0000","lat":"47.3259809",
"lng":"-119.5496231"}
*/
