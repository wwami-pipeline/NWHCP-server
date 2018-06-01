package stores

import (
	"fmt"
	"log"

	"github.com/pipeline-db/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SchoolStore struct {
	//the mongo session
	session *mgo.Session
	//the database name to use
	dbname string
	//the collection name to use
	colname string
	//the Collection object for that dbname/colname
	col *mgo.Collection
}

// Creates new School Store with mongo session, dbname, and collection name
func NewSchoolStore(sess *mgo.Session, dbName string, collectionName string) (*SchoolStore, error) {
	if sess == nil {
		panic("nil pointer passed for session")
	}

	//return a new MongoStore
	ss := &SchoolStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
		col:     sess.DB(dbName).C(collectionName),
	}

	return ss, nil
}

// Returns a school based on the schoolID
func (ss *SchoolStore) GetByID(schoolID bson.ObjectId) (*models.School, error) {
	school := &models.School{}
	if err := ss.col.FindId(schoolID).One(school); err != nil {
		return nil, nil
	}
	return school, nil
}

// Inserts a school
func (ss *SchoolStore) Insert(school *models.School) (*models.School, error) {
	log.Printf("School: %s %v", school.SchoolName)
	checkSchool, _ := ss.GetByName(school.SchoolName)
	if checkSchool != nil {
		log.Printf("School already exists, check if you want to update instead")
		return nil, nil
	} else {
		school.SchoolID = bson.NewObjectId()
		if err := ss.col.Insert(school); err != nil {
			log.Printf(err.Error())
			return nil, err
		}
		return school, nil
	}
}

// Returns a school based on the schoolName
func (ss *SchoolStore) GetByName(schoolName string) (*models.School, error) {
	school := &models.School{}
	err := ss.col.Find(bson.M{"schoolname": schoolName}).One(school)
	if err != nil {
		return nil, err
	}
	return school, nil
}

// Updates a school based on the ID
func (ss *SchoolStore) Update(schoolName string, updateSchool *models.UpdateSchool) error {
	if err := ss.col.Update(bson.M{"schoolname": schoolName}, bson.M{"$set": updateSchool}); err != nil {
		return fmt.Errorf("error updating tag: %v", err)
	}
	return nil
}

// Returns all schools in database
func (ss *SchoolStore) GetAll() ([]*models.School, error) {
	allSchools := []*models.School{}
	err := ss.col.Find(nil).All(&allSchools)
	if err != nil {
		return nil, err
	}
	return allSchools, nil
}
