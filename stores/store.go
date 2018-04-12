package stores

import (
	"pipeline-db/models"

	"gopkg.in/mgo.v2/bson"
)

//Store represents a session data store.
//This is an abstract interface that can be implemented
//against several different types of data stores. For example,
//session data could be stored in memory in a concurrent map,
//or more typically in a shared key/value server store like redis.
type Store interface {
	//GetByID returns the User with the given ID
	GetByID(id bson.ObjectId) (*models.School, error)

	GetBySchoolName(schoolName string) (*models.School, error)

	//Insert converts new tags into tags and adds it to an image.
	InsertSchool(school *models.School) (*models.School, error)

	// Delete deletes the tags associated with the tagID
	DeleteSchool(tagID bson.ObjectId) error
}
