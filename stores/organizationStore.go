package stores

import (
	"errors"
	"log"

	"pipeline-db/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
func (os *OrgStore) GetByID(orgID int) (*models.Organization, error) {
	org := &models.Organization{}

	err := os.col.Find(bson.M{"orgid": orgID}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// GetByName returns an organization based on the orgTitle
func (os *OrgStore) GetByName(orgTitle string) (*models.Organization, error) {
	org := &models.Organization{}
	err := os.col.Find(bson.M{"orgtitle": orgTitle}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

//Insert inserts an organization checks for duplicates
func (os *OrgStore) Insert(org *models.Organization) (*models.Organization, error) {
	checkOrg, _ := os.GetByName(org.OrgTitle)
	if checkOrg != nil {
		log.Printf("Organization '%s' already exists, check if you want to update instead", org.OrgTitle)
		return nil, errors.New("Organization already exists")
	}
	if err := os.col.Insert(org); err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	insertedOrg, err := os.GetByName(org.OrgTitle)
	if err != nil {
		log.Printf("Error getting the newly-inserted organization from database")
	}
	return insertedOrg, nil
}

// Update updates an organization based on the ID
func (os *OrgStore) Update(orgTitle string, updateOrg *models.Organization) (*models.Organization, error) {
	if err := os.col.Update(bson.M{"orgtitle": orgTitle}, bson.M{"$set": updateOrg}); err != nil {
		return nil, err
	}
	updatedOrg, err := os.GetByName(orgTitle)
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

// GetAll returns all organizations in database
func (os *OrgStore) GetAll() ([]*models.Organization, error) {
	allOrgs := []*models.Organization{}
	err := os.col.Find(nil).All(&allOrgs)
	if err != nil {
		return nil, err
	}
	return allOrgs, nil
}
