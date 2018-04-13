package models

import (
	"gopkg.in/mgo.v2/bson"
)

type School struct {
	SchoolID           bson.ObjectId `json:"schoolID" bson:"_id"`
	SchoolName         string        `json:"school_name"`
	SchoolDistrictName string        `json:"school_district_name"`
	FullAddress        string        `json: "full_Address"`
	Street             string        `json: "street"`
	City               string        `json: "city"`
	County             string        `json: "county"`
	Zip                string        `json: "zip"`
	Latitude           float64       `json: "lat"`
	Longitude          float64       `json: "long"`
}

type UpdateSchool struct {
	SchoolName         string  `json:"school_name"`
	SchoolDistrictName string  `json:"school_district_name"`
	FullAddress        string  `json: "full_Address"`
	Street             string  `json: "street"`
	City               string  `json: "city"`
	County             string  `json: "county"`
	Zip                string  `json: "zip"`
	Latitude           float64 `json: "lat"`
	Longitude          float64 `json: "long"`
}

/*[{"school_name":"Beezley Springs Elementary","school_district_name":"Ephrata School District",
"full_address":"501 C ST NW  EPHRATA Washington 98823-0000","street":"501 C ST NW","city":
"EPHRATA","county":"Grant County","state":"Washington","zip":"98823-0000","lat":"47.3259809",
"lng":"-119.5496231"}
*/
