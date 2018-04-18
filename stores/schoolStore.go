package stores

import (
	"fmt"
	"log"
	"reflect"

	"pipeline-db/models"

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

func (ss *SchoolStore) GetByID(schoolID bson.ObjectId) (*models.School, error) {
	school := &models.School{}
	log.Printf("schoolID: %v", schoolID)
	log.Printf("getbyid typeof: %v", reflect.TypeOf(schoolID))
	if err := ss.col.FindId(schoolID).One(school); err != nil {
		return nil, nil
	}
	log.Printf("ASDFSAFDSF: %v", school)

	return school, nil
}

func (ss *SchoolStore) InsertSchool(school *models.School) (*models.School, error) {
	log.Printf("School: %s %v", school.SchoolName)
	school.SchoolID = bson.NewObjectId()
	if err := ss.col.Insert(school); err != nil {
		log.Printf("error inserting tag")
		log.Printf(err.Error())
		return nil, err
	}
	return school, nil
}

func (ss *SchoolStore) GetBySchoolName(schoolName string) (*models.School, error) {
	school := &models.School{}
	err := ss.col.Find(bson.M{"schoolname": schoolName}).One(school)
	if err != nil {
		return nil, err
	}
	return school, nil
}

func (ss *SchoolStore) UpdateSchool(schoolName string, updateSchool *models.UpdateSchool) error {
	if err := ss.col.Update(bson.M{"schoolname": schoolName}, bson.M{"$set": updateSchool}); err != nil {
		return fmt.Errorf("error updating tag: %v", err)
	}
	return nil
}

func (ss *SchoolStore) DeleteSchool(schoolID bson.ObjectId) error {
	err := ss.col.RemoveId(schoolID)
	if err != nil {
		return err
	}
	return nil
}
