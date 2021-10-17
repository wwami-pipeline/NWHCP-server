package orgs

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	session *mongo.Client
	//the database name to use
	dbname string
	//the collection name to use
	colname string
	//the Collection object for that dbname/colname
	col *mongo.Collection
}

// NewOrgStore creates new Organization Store with mongo session, dbname, and collection name
func NewOrgStore(sess *mongo.Client, dbName string, collectionName string) (*OrgStore, error) {
	if sess == nil {
		panic("nil pointer passed for session")
	}
	//return a new MongoStore
	os := &OrgStore{
		session: sess,
		dbname:  dbName,
		colname: collectionName,
		col:     sess.Database(dbName).Collection(collectionName),
	}
	return os, nil
}

// GetByID returns an organization based on the ID
func (os *OrgStore) GetByID(orgID int) (*Organization, error) {
	org := &Organization{}
	err := os.col.FindOne(context.TODO(), bson.M{"OrgId": orgID}).Decode(org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// GetByName returns an organization based on the orgTitle
func (os *OrgStore) GetByName(orgTitle string) (*Organization, error) {
	org := &Organization{}
	err := os.col.FindOne(context.TODO(), bson.M{"OrgTitle": orgTitle}).Decode(org)
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
	_, err := os.col.InsertOne(context.TODO(), &org)
	if err != nil {
		panic(err)
	}
	insertedOrg, err := os.GetByID(org.OrgId)
	if err != nil {
		log.Printf("Error getting the newly-inserted organization from database")
	}
	return insertedOrg, nil
}

// Update updates an organization based on the ID
func (os *OrgStore) Update(orgID int, updateOrg *Organization) (*Organization, error) {
	_, err := os.col.UpdateOne(context.TODO(), bson.M{"_id": orgID}, bson.M{"$set": updateOrg})
	if err != nil {
		panic(err)
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
	_, err := os.col.DeleteOne(context.TODO(), orgID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAll truncates an organization collection
func (os *OrgStore) DeleteAll() error {
	_, err := os.col.DeleteMany(context.TODO(), bson.D{})
	return err
}

// GetAll returns all organizations in database
func (os *OrgStore) GetAll() ([]*Organization, error) {
	allOrgs := []*Organization{}
	findOptions := options.Find()
	// Sort by `OrgId` field ascending
	findOptions.SetSort(bson.D{{"OrgId", 1}})
	cursor, err := os.col.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	// collect results
	for cursor.Next(context.TODO()) {
		var org *Organization
		{
		}
		if err = cursor.Decode(&org); err != nil {
			return nil, err
		} else {
			allOrgs = append(allOrgs, org)
		}
	}
	// output
	return allOrgs, nil
}

// SearchOrgs gets the organizations that matched certain searching criteria
func (os *OrgStore) SearchOrgs(orginfo *OrgInfo) ([]*Organization, error) {
	allOrgs := []*Organization{}
	query := []bson.M{}

	if orginfo.HasCost {
		query = append(query, bson.M{
			// true false checkboxes
			"HasCost": orginfo.HasCost,
		})
	}

	if orginfo.Under18 {
		query = append(query, bson.M{
			// true false checkboxes
			"Under18": orginfo.Under18,
		})
	}

	if orginfo.HasTransport {
		query = append(query, bson.M{
			// true false checkboxes
			"HasTransport": orginfo.HasTransport,
		})
	}

	if orginfo.HasShadow {
		query = append(query, bson.M{
			// true false checkboxes
			"HasShadow": orginfo.HasShadow,
		})
	}

	// the search word is in org title street addr city state
	if len(orginfo.SearchContent) > 0 {
		query = append(query, bson.M{
			"$or": bson.A{
				bson.M{"OrgTitle": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"StreetAddr": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"City": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"State": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"ZipCode": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"Email": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"ActivityDesc": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
				bson.M{"CareerEmp": bson.M{"$regex": "(?i)" + orginfo.SearchContent}},
			},
		})
	}
	// career selection
	if len(orginfo.CareerEmp) > 0 {
		query = append(query, bson.M{
			"CareerEmp": bson.M{"$all": orginfo.CareerEmp},
		})
	}
	// grade levels
	if len(orginfo.GradeLevels) > 0 {
		query = append(query, bson.M{
			"GradeLevels": bson.M{"$all": orginfo.GradeLevels},
		})
	}

	// query
	inputQuery := bson.M{}
	if len(query) > 0 {
		inputQuery = bson.M{"$and": query}
	}
	findOptions := options.Find()
	// Sort by `OrgId` field ascending
	findOptions.SetSort(bson.D{{"OrgId", 1}})
	cursor, err := os.col.Find(context.TODO(), inputQuery, findOptions)
	if err != nil {
		return nil, err
	}
	// collect results
	for cursor.Next(context.Background()) {
		var org *Organization
		{
		}
		if err = cursor.Decode(&org); err != nil {
			return nil, err
		} else {
			allOrgs = append(allOrgs, org)
		}
	}
	// output
	return allOrgs, nil
}
