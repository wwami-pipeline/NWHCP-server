package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"pipeline-db/models/orgs"
	"strconv"
	"strings"
	// "../nwhcp/nwhcp-server/gateway/models/users"
)

const contentTypeHeader = "Content-Type"
const contentTypeApplicationJSON = "application/json"

// func (ctx *HandlerContext) InsertOrgs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add(contentTypeHeader, contentTypeApplicationJSON) // if os.Getenv("APP_ENV") == "production" && r.Header.Get("AUTH_TOKEN_FOR_PYTHON") != os.Getenv("AUTH_TOKEN_FOR_PYTHON") {
// 	// 	return
// 	// }
// 	switch r.Method {
// 	case "POST":
// 		if !strings.HasPrefix(r.Header.Get(contentTypeHeader), contentTypeApplicationJSON) {
// 			http.Error(w, fmt.Sprintf("The request body must be in JSON"), http.StatusUnsupportedMediaType)
// 			return
// 		}
// 		var o []orgs.Organization

// 		if err := json.NewDecoder(r.Body).Decode(&o); err != nil {
// 			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}

// 		var insertedOrgs []orgs.Organization

// 		for _, org := range o {
// 			dbOrg := &orgs.Organization{}
// 			dbOrg.OrgId = org.OrgId
// 			dbOrg.OrgTitle = org.OrgTitle
// 			dbOrg.OrgWebsite = org.OrgWebsite
// 			dbOrg.StreetAddress = org.StreetAddress
// 			dbOrg.City = org.City
// 			dbOrg.State = org.State
// 			dbOrg.ZipCode = org.ZipCode
// 			dbOrg.Phone = org.Phone
// 			dbOrg.Email = org.Email
// 			dbOrg.ActivityDesc = org.ActivityDesc
// 			dbOrg.Lat = org.Lat
// 			dbOrg.Long = org.Long
// 			dbOrg.HasShadow = org.HasShadow
// 			dbOrg.HasCost = org.HasCost
// 			dbOrg.HasTransport = org.HasTransport
// 			dbOrg.Under18 = org.Under18
// 			dbOrg.CareerEmp = org.CareerEmp
// 			dbOrg.GradeLevels = org.GradeLevels
// 			insertedOrg, err := ctx.OrgStore.Insert(dbOrg)
// 			if err != nil {
// 				http.Error(w, fmt.Sprintf("Error inserting new organization '%v' into the database: %v", org.OrgTitle, err),
// 					http.StatusBadRequest)
// 			} else {
// 				insertedOrgs = append(insertedOrgs, *insertedOrg)
// 			}
// 			insq := "INSERT INTO organization(org_title) VALUES(?)"
// 			_, err2 := ctx.Db.Exec(insq, org.OrgTitle)
// 			if err2 != nil {
// 				http.Error(w, "Error with connecting to database", http.StatusInternalServerError)
// 				return
// 			}
// 		}
// 		w.WriteHeader(http.StatusCreated)

// 		if err := json.NewEncoder(w).Encode(insertedOrgs); err != nil {
// 			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
// 				http.StatusInternalServerError)
// 			return
// 		}

// 	default:
// 		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// //GetAllOrgs is used to grab organization data
// func (ctx *HandlerContext) GetAllOrgs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)
// 	log.Println("hello getting all orgs")
// 	switch r.Method {
// 	case "GET":
// 		allOrgs, err := ctx.OrgStore.GetAll()
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error getting organizations fromn the data store:  %v", err),
// 				http.StatusNotFound)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)

// 		if err := json.NewEncoder(w).Encode(allOrgs); err != nil {
// 			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
// 				http.StatusInternalServerError)
// 			return
// 		}

// 	default:
// 		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// func (ctx *HandlerContext) DeleteAllOrgsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodDelete:
// 		err := ctx.OrgStore.DeleteAll()
// 		if err != nil {
// 			http.Error(w, "Cannot truncate collection", http.StatusBadRequest)
// 			return
// 		}
// 	default:
// 		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// SpecificOrgHandler handles requests for a specific organization.
// The resource path will be /api/v1/org/{id}}, where {id} will be the organization's ID.
func (ctx *HandlerContext) SpecificOrgHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("hello getting specific orgs")
	idString := path.Base(r.URL.Path)

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, fmt.Sprintf("Did not provide {id} as a number in /v1/org/{id}, please provide the correct ID"),
			http.StatusForbidden)
		return
	}

	org, err := ctx.OrgStore.GetByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("No organization is found with the ID %d in the data store:  %v", id, err),
			http.StatusNotFound)
		return
	}

	switch r.Method {

	case http.MethodGet:
		w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(org); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}

	case http.MethodPatch:
		if !strings.HasPrefix(r.Header.Get(contentTypeHeader), contentTypeApplicationJSON) {
			http.Error(w, fmt.Sprintf("The request body must be in JSON"), http.StatusUnsupportedMediaType)
			return
		}

		var updateOrg orgs.Organization

		if err := json.NewDecoder(r.Body).Decode(&updateOrg); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
				http.StatusBadRequest)
			return
		}

		updatedOrg, err := ctx.OrgStore.Update(org.OrgId, &updateOrg)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error updating the organization: %v", err),
				http.StatusBadRequest)
			return
		}

		w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(updatedOrg); err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
				http.StatusInternalServerError)
			return
		}

	case http.MethodDelete:
		err := ctx.OrgStore.Delete(org.OrgId)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error deleting the organization: %v", err),
				http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The orgainzation was successful deleted"))

	case http.MethodPost:
		err := CheckAuth(w, r, ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error user is not authenticated: %v", err), http.StatusUnauthorized)
			return
		}
		insertQ := "INSERT INTO user_org(UserID, OrgID) VALUES(?,?)"
		_, insertErr := ctx.Db.Exec(insertQ, ctx.UserID, id)
		if insertErr != nil {
			fmt.Printf(": %v", insertErr)
			http.Error(w, "Error inserting into database", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Organization successfully added to user account"))

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}

// func (ctx *HandlerContext) SearchOrgsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		if !strings.HasPrefix(r.Header.Get(contentTypeHeader), contentTypeApplicationJSON) {
// 			http.Error(w, fmt.Sprintf("The request body must be in JSON"), http.StatusUnsupportedMediaType)
// 			return
// 		}

// 		orgInfo := &orgs.OrgInfo{}
// 		if err := json.NewDecoder(r.Body).Decode(orgInfo); err != nil {
// 			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}

// 		orgs, err := ctx.OrgStore.SearchOrgs(orgInfo)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error getting the org info from the database: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}

// 		w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)
// 		w.WriteHeader(http.StatusOK)

// 		if err := json.NewEncoder(w).Encode(orgs); err != nil {
// 			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
// 				http.StatusInternalServerError)
// 			return
// 		}

// 	default:
// 		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// GetUserInfoHandler blah
// add check auth to get the user's info
func (ctx *HandlerContext) GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method must be GET", http.StatusMethodNotAllowed)
		return
	}
	err := CheckAuth(w, r, ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error user is not authenticated: %v", err), http.StatusUnauthorized)
		return
	}
	userOrgs := []UserOrgs{}
	log.Println(ctx.UserID)
	query := "SELECT UO.UserID, UO.OrgID FROM user_org UO WHERE UO.UserID = ?"
	rows, qErr := ctx.Db.Query(query, ctx.UserID)
	if qErr != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		curRow := UserOrgs{}
		if err := rows.Scan(&curRow.UserID, &curRow.OrgID); err != nil {
			http.Error(w, "Error getting User Orgs", http.StatusInternalServerError)
			return
		}
		curOrg, err := ctx.OrgStore.GetByID(int(curRow.OrgID))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error getting organization: %v", err), http.StatusInternalServerError)
			return
		}
		curRow.OrgTitle = curOrg.OrgTitle
		curRow.OrgWebsite = curOrg.OrgWebsite
		curRow.OrgCity = curOrg.City
		curRow.OrgState = curOrg.State
		curRow.OrgZipcode = curOrg.ZipCode
		curRow.OrgPhone = curOrg.Phone
		userOrgs = append(userOrgs, curRow)
	}
	// delete orgs table just user orgs
	json, jErr := json.Marshal(userOrgs)
	if jErr != nil {
		http.Error(w, "Data could not be returned", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(json)
}

func CheckAuth(w http.ResponseWriter, r *http.Request, c *HandlerContext) error {
	log.Println(r.Header.Get("X-User"))
	if r.Header.Get("X-User") == "" {
		http.Error(w, "User is not authenticated", http.StatusUnauthorized)
		return errors.New("unauthenticated user")
	}
	userInfo := &Creds{}
	json.Unmarshal([]byte(r.Header.Get("X-User")), userInfo)
	log.Println(userInfo)
	c.UserID = userInfo.ID
	return nil
}
