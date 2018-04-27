package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pipeline-db/models"
)

func (ctx *Context) PopulateDB(w http.ResponseWriter, r *http.Request) {
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
			ctx.Store.InsertSchool(dbSchool)

		}
	}
}

func (ctx *Context) SchoolHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		allSchools, _ := ctx.Store.GetAll()
		err := json.NewEncoder(w).Encode(allSchools)
		if err != nil {

			http.Error(w, "Unable to encode json", http.StatusInternalServerError)
			return
		}

	}
}
