package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	// "../nwhcp/nwhcp-server/gateway/models/users"
)

const contentTypeHeader = "Content-Type"
const contentTypeApplicationJSON = "application/json"

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

	_, getErr := ctx.OrgStore.GetByID(id)
	if getErr != nil {
		http.Error(w, fmt.Sprintf("No organization is found with the ID %d in the data store:  %v", id, err),
			http.StatusNotFound)
		return
	}

	authErr := CheckAuth(w, r, ctx)
	if authErr != nil {
		http.Error(w, fmt.Sprintf("Error user is not authenticated: %v", authErr), http.StatusUnauthorized)
		return
	}

	switch r.Method {

	case http.MethodDelete:
		deleteQ := "DELETE FROM user_org WHERE UserID = ? AND OrgID = ?"
		_, delErr := ctx.Db.Exec(deleteQ, ctx.UserID, id)
		if delErr != nil {
			http.Error(w, "Error deleting org from account", http.StatusInternalServerError)
			return
		}
		w.Header().Set(contentTypeHeader, "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The orgainzation was successful deleted from the user's account"))

	case http.MethodPost:

		insertQ := "INSERT INTO user_org(UserID, OrgID) VALUES(?,?)"
		_, insertErr := ctx.Db.Exec(insertQ, ctx.UserID, id)
		if insertErr != nil {
			fmt.Printf(": %v", insertErr)
			http.Error(w, "Error inserting into database", http.StatusInternalServerError)
			return
		}
		w.Header().Set(contentTypeHeader, "text/html")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Organization successfully added to user account"))

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}

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
	query := "SELECT UserID, OrgID FROM user_org UO WHERE UserID = ? GROUP BY UserID, OrgID"
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
	w.Header().Set(contentTypeHeader, contentTypeApplicationJSON)
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
