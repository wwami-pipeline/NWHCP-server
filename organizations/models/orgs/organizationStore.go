package orgs

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//These are the fields that being searched by client
const (
	HasShadow    = "HasShadow"
	HasCost      = "HasCost"
	HasTransport = "HasTransport"
	Under18      = "Under18"
	CareerEmp    = "CareerEmp"
	GradeLevels  = "GradeLevels"
	OrgTitle     = "OrgTitle"
	StreetAddr   = "StreetAddress"
	City         = "City"
	State        = "State"
)

//OrgStore represents a mongoDB data store that implements the abstract store interface
type OrgStore struct {
	//the mongo session
	session *mgo.Session
	//the database name to use
	dbname string
	//the collection name to use
	colname string
	//the Collection object for that dbname/colname
	col *mgo.Collection
}

// NewOrgStore creates new Organization Store with mongo session, dbname, and collection name
func NewOrgStore(sess *mgo.Session, dbName string, collectionName string) (*OrgStore, error) {
	if sess == nil {
		panic("nil pointer passed for session")
	}

	//return a new MongoStore
	os := &OrgStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
		col:     sess.DB(dbName).C(collectionName),
	}

	return os, nil
}

// GetByID returns an organization based on the ID
func (os *OrgStore) GetByID(orgID int) (*Organization, error) {
	org := &Organization{}
	err := os.col.Find(bson.M{"_id": orgID}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// GetByName returns an organization based on the orgTitle
func (os *OrgStore) GetByName(orgTitle string) (*Organization, error) {
	org := &Organization{}
	err := os.col.Find(bson.M{"OrgTitle": orgTitle}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

//Insert inserts an organization checks for duplicates
func (os *OrgStore) Insert(org *Organization) (*Organization, error) {
	checkOrg, _ := os.GetByID(org.OrgId)
	if checkOrg != nil {
		log.Printf("Organization with ID '%s' already exists, updating", org.OrgId)
		return nil, errors.New("Organization already exists")
	}
	if err := os.col.Insert(org); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	insertedOrg, err := os.GetByID(org.OrgId)
	if err != nil {
		log.Printf("Error getting the newly-inserted organization from database")
	}
	return insertedOrg, nil
}

// Update updates an organization based on the ID
func (os *OrgStore) Update(orgID int, updateOrg *Organization) (*Organization, error) {
	if err := os.col.Update(bson.M{"_id": orgID}, bson.M{"$set": updateOrg}); err != nil {
		return nil, err
	}
	updatedOrg, err := os.GetByID(orgID)
	if err != nil {
		log.Printf("Error getting the updated organization from the database: %v\n", err)
		return nil, err
	}
	return updatedOrg, nil
}

// Delete deletes an organization based on the ID
func (os *OrgStore) Delete(orgID int) error {
	err := os.col.RemoveId(orgID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll truncates an organization collection
func (os *OrgStore) DeleteAll() error {
	_, err := os.col.RemoveAll(nil)
	return err
}

// GetAll returns all organizations in database
func (os *OrgStore) GetAll() ([]*Organization, error) {
	allOrgs := []*Organization{}
	err := os.col.Find(nil).All(&allOrgs)
	if err != nil {
		return nil, err
	}
	return allOrgs, nil
}

// SearchOrgs gets the organizations that matched certain searching criteria
func (os *OrgStore) SearchOrgs(orginfo *OrgInfo) ([]*Organization, error) {
	allOrgs := []*Organization{}

	andQuery := []bson.M{}
	andQuery = append(andQuery,
		bson.M{"$or": buildOrQueryForSearchContent(orginfo)},
		bson.M{"$or": buildOrQueryForCareerEmp(orginfo)},
		bson.M{"$or": buildOrQueryForGradeLevels(orginfo)},
		bson.M{"$and": andQueryForCheckBox(orginfo)})

	err := os.col.Find(bson.M{"$and": andQuery}).All(&allOrgs)
	if err != nil {
		return nil, err
	}

	return allOrgs, nil
}

func buildOrQueryForSearchContent(orginfo *OrgInfo) []bson.M {
	searchFields := make([]string, 4)
	searchFields = append(searchFields, OrgTitle, StreetAddr, City, State)

	orQuery := []bson.M{}
	for _, field := range searchFields {
		query := bson.M{field: bson.M{"$regex": orginfo.SearchContent, "$options": "i"}}
		orQuery = append(orQuery, query)
	}
	if len(orQuery) == 0 {
		orQuery = append(orQuery, nil)
	}
	return orQuery
}

func buildOrQueryForCareerEmp(orginfo *OrgInfo) []bson.M {
	orQuery := []bson.M{}
	for _, c := range orginfo.CareerEmp {
		query := bson.M{CareerEmp: c}
		orQuery = append(orQuery, query)
	}
	if len(orQuery) == 0 {
		orQuery = append(orQuery, nil)
	}
	return orQuery
}

func buildOrQueryForGradeLevels(orginfo *OrgInfo) []bson.M {
	orQuery := []bson.M{}

	for _, c := range orginfo.GradeLevels {
		query := bson.M{GradeLevels: c}
		orQuery = append(orQuery, query)
	}
	if len(orQuery) == 0 {
		orQuery = append(orQuery, nil)
	}
	return orQuery
}

func andQueryForCheckBox(orginfo *OrgInfo) []bson.M {
	andQuery := []bson.M{}
	if orginfo.HasShadow {
		query := bson.M{HasShadow: orginfo.HasShadow}
		andQuery = append(andQuery, query)
	}
	if orginfo.HasCost {
		query := bson.M{HasCost: orginfo.HasCost}
		andQuery = append(andQuery, query)
	}
	if orginfo.HasTransport {
		query := bson.M{HasTransport: orginfo.HasTransport}
		andQuery = append(andQuery, query)
	}
	if orginfo.Under18 {
		query := bson.M{Under18: orginfo.Under18}
		andQuery = append(andQuery, query)
	}
	if len(andQuery) == 0 {
		andQuery = append(andQuery, nil)
	}

	return andQuery
}
