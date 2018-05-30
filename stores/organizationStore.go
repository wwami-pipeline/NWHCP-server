package stores

import (
	"fmt"
	"log"

	"github.com/pipeline-db/models"
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

// func (os *OrganizationStore) GetByID(orgID bson.ObjectId) (*models.Organization, error) {
// 	org := &models.Organization{}
// 	if err := os.col.FindId(orgID).One(org); err != nil {
// 		return nil, nil
// 	}
// 	return org, nil
// }
func (os *OrganizationStore) GetByID(orgID int) (*models.Organization, error) {
	org := &models.Organization{}
	err := os.col.Find(bson.M{"orgID": orgID}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (os *OrganizationStore) Insert(org *models.Organization) (*models.Organization, error) {
	log.Printf("Organization: %s %v", org.OrgTitle)
	checkOrg, _ := os.GetByName(org.OrgTitle)
	if checkOrg != nil {
		log.Printf("Organization already exists, check if you want to update instead")
		return nil, nil
	} else {
		// org.OrganizationId = bson.NewObjectId()
		if err := os.col.Insert(org); err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		return org, nil
	}
}

func (os *OrganizationStore) GetByName(orgTitle string) (*models.Organization, error) {
	org := &models.Organization{}
	err := os.col.Find(bson.M{"OrgTitle": orgTitle}).One(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (os *OrganizationStore) Update(orgName string, updateOrg *models.UpdateOrganization) error {
	if err := os.col.Update(bson.M{"orgname": orgName}, bson.M{"$set": updateOrg}); err != nil {
		return fmt.Errorf("error updating organization: %v", err)
	}
	return nil
}

func (os *SchoolStore) Delete(orgID bson.ObjectId) error {
	err := os.col.RemoveId(orgID)
	if err != nil {
		return err
	}
	return nil
}

func (os *OrganizationStore) GetAll() ([]*models.Organization, error) {
	allOrgs := []*models.Organization{}
	err := os.col.Find(nil).All(&allOrgs)
	if err != nil {
		return nil, err
	}
	return allOrgs, nil
}
