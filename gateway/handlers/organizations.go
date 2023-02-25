package handlers

// controllers for orgs- not in use but refactoring
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"nwhcp/nwhcp-server/gateway/models/orgs"
	"os"

	// "pipeline-db/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const contentTypeHeader = "Content-Type"
const contentTypeApplicationJSON = "application/json"

// assign type to all funcs for easy access in main
// User controller should have access to a Mongo session
type OrganizationController struct {
	session *mongo.Client
}

// convenience function returns UserController
// pass around the address, not the big data structure
func NewOrganizationController(s *mongo.Client) *OrganizationController {
	return &OrganizationController{s}
}

func (oc OrganizationController) CreateOrganization(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	o := &orgs.NewOrganization{
		OrgTitle:      params["orgTitle"],
		OrgWebsite:    params["orgWebsite"],
		StreetAddress: params["streetAddress"],
		City:          params["city"],
		State:         params["state"],
		ZipCode:       params["zipCode"],
		Phone:         params["phone"],
		Email:         params["email"],
	}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&o)

	// Create BSON ID
	o.ID = primitive.NewObjectID()

	// does this create a collection if it doesn't exist?
	oc.session.Database("mongodb").Collection("organizations").InsertOne(context.TODO(), &o)

	json.NewEncoder(w).Encode(&o)

}

// // InsertOrgs inserts organization data
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

// 		var orgsList []orgs.Organization

// 		if err := json.NewDecoder(r.Body).Decode(&orgsList); err != nil {
// 			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}

// 		var insertedOrgs []orgs.Organization

// 		for _, org := range orgsList {
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

// // GetAllOrgs is used to grab organization data
// func (ctx *HandlerContext) GetAllOrgs(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)

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

// // SpecificOrgHandler handles requests for a specific organization.
// The resource path will be /api/v1/org/{id}}, where {id} will be the organization's ID.
// receiver before - ctx *HandlerContext
func (oc OrganizationController) GetOrgByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "applications/json")

	params := mux.Vars(r)

	id := params["id"]

	oid, _ := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	fmt.Println("ObjectIDFromHex ERROR:", err)
	// }

	// empty struct to store org
	result := &orgs.Organization{}

	// Fetch org
	err := oc.session.Database("mongodb").Collection("surveys").FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		fmt.Println("Organization not found")
		os.Exit(1)
		return
	}

	json.NewEncoder(w).Encode(*result)

}

// idString := path.Base(r.URL.Path)

// id, err := strconv.Atoi(idString)
// if err != nil {
// 	http.Error(w, fmt.Sprintf("Did not provide {id} as a number in /v1/org/{id}, please provide the correct ID"),
// 		http.StatusForbidden)
// 	return
// }

// org, err := ctx.OrgStore.GetByID(id)
// if err != nil {
// 	http.Error(w, fmt.Sprintf("No organization is found with the ID %d in the data store:  %v", id, err),
// 		http.StatusNotFound)
// 	return
// }

// 	switch r.Method {

// 	case http.MethodGet:
// 		w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)
// 		w.WriteHeader(http.StatusOK)
// 		if err := json.NewEncoder(w).Encode(org); err != nil {
// 			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
// 				http.StatusInternalServerError)
// 			return
// 		}

// 	case http.MethodPatch:
// 		if !strings.HasPrefix(r.Header.Get(contentTypeHeader), contentTypeApplicationJSON) {
// 			http.Error(w, fmt.Sprintf("The request body must be in JSON"), http.StatusUnsupportedMediaType)
// 			return
// 		}

// 		var updateOrg orgs.Organization

// 		if err := json.NewDecoder(r.Body).Decode(&updateOrg); err != nil {
// 			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}

// 		updatedOrg, err := ctx.OrgStore.Update(org.OrgId, &updateOrg)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error updating the organization: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}

// 		w.Header().Add(contentTypeHeader, contentTypeApplicationJSON)
// 		w.WriteHeader(http.StatusOK)

// 		if err := json.NewEncoder(w).Encode(updatedOrg); err != nil {
// 			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err),
// 				http.StatusInternalServerError)
// 			return
// 		}

// 	case http.MethodDelete:
// 		err := ctx.OrgStore.Delete(org.OrgId)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error deleting the organization: %v", err),
// 				http.StatusBadRequest)
// 			return
// 		}
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("The orgainzation was successful deleted"))

// 	default:
// 		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
// 		return
// 	}
// }

// func (ctx *HandlerContext) SearchOrgsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		if !strings.HasPrefix(r.Header.Get(contentTypeHeader), contentTypeApplicationJSON) {
// 			http.Error(w, fmt.Sprintf("The request body must be in JSON"), http.StatusUnsupportedMediaType)
// 			return
// 		}

// 		// debug
// 		// Save a copy of this request for debugging.
// 		//requestDump, err := httputil.DumpRequest(r, true)
// 		//if err != nil {
// 		//	fmt.Println(err)
// 		//}
// 		//fmt.Println(string(requestDump))
// 		// end debug

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
