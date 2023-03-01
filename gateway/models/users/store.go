package models

import (
	"errors"
	"time"
)

// ErrUserNotFound is returned when the user can't be found
var ErrUserNotFound = errors.New("user not found")

// Store represents a store for Users
type UserStore interface {
	//GetByID returns the User with the given ID
	GetByID(id int64) (*models.User, error)

	//GetByEmail returns the User with the given email
	GetByEmail(email string) (*models.User, error)

	//Insert inserts the user into the database, and returns
	//the newly-inserted User, complete with the DBMS-assigned ID
	Insert(user *models.User) (*models.User, error)

	//Update applies UserUpdates to the given user ID
	//and returns the newly-updated user
	Update(id int64, updates *models.Updates) (*models.User, error)

	//Delete deletes the user with the given ID
	Delete(id int64) error

	//Tracks a single login
	TrackLogin(id int64, ip string, time time.Time) error

	//GetUserOrgs
	GetOrgs(id int64) (*models.UserOrgs, error)
}
