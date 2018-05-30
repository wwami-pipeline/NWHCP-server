package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pipeline-db/models"
)

func (ctx *Context) PopulateDB(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)

		var schools []models.School
		err := decoder.Decode(&schools)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
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
			ctx.Store1.Insert(dbSchool)

		}
	}
}

func (ctx *Context) PopulateDB2(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)

		var orgs []models.Organization
		err := decoder.Decode(&orgs)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		for _, org := range orgs {
			dbOrg := &models.Organization{}
			dbOrg.OrgId = org.OrgId
			print(org.OrgId)
			dbOrg.OrgTitle = org.OrgTitle
			dbOrg.OrgWebsite = org.OrgWebsite
			dbOrg.StreetAddress = org.StreetAddress
			dbOrg.City = org.City

			dbOrg.State = org.State
			dbOrg.ZipCode = org.ZipCode
			dbOrg.Phone = org.Phone
			dbOrg.Email = org.Email
			dbOrg.ActivityDesc = org.ActivityDesc
			dbOrg.Lat = org.Lat
			dbOrg.Long = org.Long
			dbOrg.HasShadow = org.HasShadow
			dbOrg.HasCost = org.HasCost
			dbOrg.HasTransport = org.HasTransport
			dbOrg.Under18 = org.Under18
			dbOrg.CareerEmp = org.CareerEmp
			dbOrg.GradeLevels = org.GradeLevels
			ctx.Store2.Insert(dbOrg)
		}
	}
}

func (ctx *Context) SchoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	switch r.Method {
	case "GET":
		allSchools, _ := ctx.Store1.GetAll()
		err := json.NewEncoder(w).Encode(allSchools)
		if err != nil {

			http.Error(w, "Unable to encode json", http.StatusInternalServerError)
			return
		}

	}
}
func (ctx *Context) OrgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		allOrgs, _ := ctx.Store2.GetAll()
		err := json.NewEncoder(w).Encode(allOrgs)
		if err != nil {

			http.Error(w, "Unable to encode json", http.StatusInternalServerError)
			return
		}

	}
}
