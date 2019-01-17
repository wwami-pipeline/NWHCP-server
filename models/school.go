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
