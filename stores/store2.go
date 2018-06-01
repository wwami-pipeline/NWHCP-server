package stores

import (
	"github.com/pipeline-db/models"
	"gopkg.in/mgo.v2/bson"
)

//Store represents a session data store.
//This is an abstract interface that can be implemented
//against several different types of data stores. For example,
//session data could be stored in memory in a concurrent map,
//or more typically in a shared key/value server store like redis.
type Store2 interface {
	//GetByID returns the User with the given ID
	GetByID(id int) (*models.Organization, error)

	GetByName(orgName string) (*models.Organization, error)

	//Insert converts new tags into tags and adds it to an image.
	Insert(org *models.Organization) (*models.Organization, error)

	// Updates the school based on the schoolname.
	Update(orgName string, updateOrganization *models.UpdateOrganization) error

	// Delete deletes the tags associated with the tagID
	Delete(orgID bson.ObjectId) error

	// Get all schools in the database
	GetAll() ([]*models.Organization, error)
}
