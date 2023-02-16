package users

import "go.mongodb.org/mongo-driver/bson/primitive"

// models

// bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

// User represents a user account in the database (student)
// how do I store the stuff?
type User struct {
	ID                     int64   `bson:"_id"`
	Email                  string  `json:"-"` //never JSON encoded/decoded
	PassHash               []byte  `json:"-"` //never JSON encoded/decoded
	FirstName              string  `json:"firstName"`
	LastName               string  `json:"lastName"`
	JoinDate               string  `json:"joinDate"`
	Gender                 string  `json:gender`
	Age                    string  `json:age`
	State                  string  `json:state`
	FavoritedOrganizations []*Org  `json:favoritedOrganizations`
	CompletedPrograms      []*Org  `json:completedPrograms`
	InProcessPrograms      []*Org  `json:inProcessPrograms`
	Notes                  []*Note `json:notes`
}

// Credentials represents user sign-in credentials (student)
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewUser represents a new user signing up for an account (student)
type NewUser struct {
	ID           primitive.ObjectID `"bson:id"`
	Email        string             `json:"email"`
	Password     string             `json:"password"`
	PasswordConf string             `json:"passwordConf"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName"`
	Age          string             `json:age`
	State        string             `json:state`
}

// Updates represents updates allowed to be edited by user (student)
type Updates struct {
	FirstName              string `json:"firstName"`
	LastName               string `json:"lastName"`
	State                  string `json:state`
	FavoritedOrganizations []*Org `json:favoritedOrganizations`
	CompletedPrograms      []*Org `json:completedPrograms`
	InProcessPrograms      []*Org `json:inProcessPrograms`
	Gender                 string `json:gender`
}

// Orgs represents the users' (student) organizations
type Org struct {
	OrgID         int64  `json:"orgId"`
	OrgTitle      string `json:"orgTitle"`
	OrgLocation   string `json:orgLocation`
	OrgType       string `json:orgType`
	OrgCompleting bool   `json:orgCompleting`
	OrgCompleted  bool   `json:orgCompleted`
}

// Notes represents users' (student) notes about programs
type Note struct {
	NoteID      int64  `json:noteId`
	NoteContent string `json:noteContent`
	OrgID       int64  `json:orgId`
	CreatedAt   string `json:createdAt`
	UpdatedAt   string `json:updatedAt`
	Reviewed    bool   `json:reviewed`
}

// UserOrgs represents a program administrator's organizations
// needed??
type UserOrgs struct {
	ID        int64  `json:"id"`
	Email     string `json:"-"` //never JSON encoded/decoded
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Orgs      []*Org `json:"orgs"`
}

// Validate validates the new user and returns an error if
// any of the validation rules fail, or nil if its valid
// func (nu *NewUser) Validate() error {
// 	_, err := mail.ParseAddress(nu.Email)
// 	if err != nil {
// 		return fmt.Errorf("Invalid Email")
// 	}
// 	if len(nu.Password) < 6 {
// 		return fmt.Errorf("Password must be at least 6 characters")
// 	}
// 	if len(nu.Password) != len(nu.PasswordConf) || strings.Compare(nu.Password, nu.PasswordConf) != 0 {
// 		return fmt.Errorf("Password and confirmation do not match")
// 	}
// 	return nil
// }

// ToUser converts the NewUser to a User, setting the
// PassHash field appropriately
// func (nu *NewUser) ToUser() (*User, error) {
// 	err := nu.Validate()
// 	if err != nil {
// 		return nil, err
// 	}
// 	user := &User{}
// 	user.FirstName = nu.FirstName
// 	user.LastName = nu.LastName
// 	user.Email = nu.Email
// 	joinDate := time.Now().Format("01-02-2006")
// 	user.JoinDate = joinDate
// 	user.SetPassword(nu.Password)

// 	return user, nil
// }

// FullName returns the user's full name, in the form:
// "<FirstName> <LastName>"
// If either first or last name is an empty string, no
// space is put between the names. If both are missing,
// this returns an empty string
// func (u *User) FullName() string {
// 	if len(u.FirstName) == 0 && len(u.LastName) == 0 {
// 		return ""
// 	} else if len(u.FirstName) == 0 {
// 		return u.LastName
// 	} else if len(u.LastName) == 0 {
// 		return u.FirstName
// 	} else {
// 		return u.FirstName + " " + u.LastName
// 	}
// }

// SetPassword hashes the password and stores it in the PassHash field
// func (u *User) SetPassword(password string) error {
// 	temp, err := bcrypt.GenerateFromPassword([]byte(password), 10)
// 	if err != nil {
// 		return err
// 	}
// 	u.PassHash = temp
// 	return nil
// }

// Authenticate compares the plaintext password against the stored hash
// and returns an error if they don't match, or nil if they do
// func (u *User) Authenticate(password string) error {
// 	return bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
// }

// // ApplyUpdates applies the updates to the user. An error
// // is returned if the updates are invalid
// func (u *User) ApplyUpdates(updates *Updates) error {
// 	if updates.FirstName == "" && updates.LastName == "" {
// 		return fmt.Errorf("names cannot both be null")
// 	}
// 	u.FirstName = updates.FirstName
// 	u.LastName = updates.LastName
// 	return nil
// }
