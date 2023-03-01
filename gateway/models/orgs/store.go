package models

// Store represents a session data store.
// This is an abstract interface that can be implemented
// against several different types of data stores. For example,
// session data could be stored in memory in a concurrent map,
// or more typically in a shared key/value server store like redis.
type OrgStore interface {
	//GetByID returns the Org with the given ID
	GetByID(id int) (*models.Organization, error)

	//GetByName returns the Org with the given game
	GetByName(orgName string) (*models.Organization, error)

	//Insert inserts the organization into database
	Insert(org *models.Organization) (*models.Organization, error)

	// Updates the organization based on the name.
	Update(orgID int, updateOrganization *models.Organization) (*models.Organization, error)

	// Delete deletes the organization associated with the ID
	Delete(orgID int) error

	// DeleteAll deletes the organization collection
	DeleteAll() error

	// Get all organizations in the database
	GetAll() ([]*models.Organization, error)

	//SearchOrgs get the organizations that matched certain searching criteria
	SearchOrgs(orginfo *models.OrgInfo) ([]*models.Organization, error)
}
