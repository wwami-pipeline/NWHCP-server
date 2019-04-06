package stores

import (
	"fmt"
	"log"

	"pipeline-db/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type OrganizationStore struct {
	//the mongo session
	session *mgo.Session
	//the database name to use
	dbname string
	//the collection name to use
	colname string
	//the Collection object for that dbname/colname
	col *mgo.Collection
}

// Creates new Organization Store with mongo session, dbname, and collection name
func NewOrganizationStore(sess *mgo.Session, dbName string, collectionName string) (*OrganizationStore, error) {
	if sess == nil {
		panic("nil pointer passed for session")
	}

	//return a new MongoStore
	os := &OrganizationStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
		col:     sess.DB(dbName).C(collectionName),
	}

	return os, nil
}

// Returns an organization based on the ID
func (os *OrganizationStore) GetByID(orgID int) (*models.Organization, error) {
	org := &models.Organization{}
	err := os.col.Find(bson.M{"orgID": orgID}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// Inserts an organization checks for duplicates
func (os *OrganizationStore) Insert(org *models.Organization) (*models.Organization, error) {
	checkOrg, _ := os.GetByName(org.OrgTitle)
	if checkOrg != nil {
		log.Printf("Organization already exists, check if you want to update instead")
		return nil, nil
	} else {
		if err := os.col.Insert(org); err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		return org, nil
	}
}

// Returns an organization based on the orgTitle
func (os *OrganizationStore) GetByName(orgTitle string) (*models.Organization, error) {
	org := &models.Organization{}
	err := os.col.Find(bson.M{"OrgTitle": orgTitle}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// Updates an organization based on the ID
func (os *OrganizationStore) Update(orgTitle string, updateOrg *models.UpdateOrganization) error {
	if err := os.col.Update(bson.M{"OrgTitle": orgTitle}, bson.M{"$set": updateOrg}); err != nil {
		return fmt.Errorf("error updating organization: %v", err)
	}
	return nil
}

// Deletes an organization based on the ID
func (os *SchoolStore) Delete(orgID bson.ObjectId) error {
	err := os.col.RemoveId(orgID)
	if err != nil {
		return err
	}
	return nil
}

// Returns all organizations in database
func (os *OrganizationStore) GetAll() ([]*models.Organization, error) {
	allOrgs := []*models.Organization{}
	err := os.col.Find(nil).All(&allOrgs)
	if err != nil {
		return nil, err
	}
	// log.Printf("Org getAll() called %v", allOrgs[0])
	log.Println(len(allOrgs))
	return allOrgs, nil
}
