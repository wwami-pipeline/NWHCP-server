package users

import (
	"nwhcp/nwhcp-server/gateway/models/orgs"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// models

// bcryptCost is the default bcrypt cost to use when hashing passwords
var bcryptCost = 13

// User represents a user account in the database (student)
// how do I store the stuff?
type User struct {
	ID                     primitive.ObjectID   `bson:"_id, omitempty"`
	Email                  string               `bson:"email" json:"email"`       //never JSON encoded/decoded
	PassHash               []byte               `bson:"password" json:"password"` //never JSON encoded/decoded
	FirstName              string               `bson: "firstName" json:"firstName"`
	JoinDate               string               `bson: "joinDate" json:"joinDate"`
	Gender                 string               `bson: "gender" json:"gender"`
	Age                    string               `bson: "age" json:"age"`
	State                  string               `bson: "state" json:"state"`
	FavoritedOrganizations []*orgs.Organization `bson: "favoritedOrganizations" json:"favoritedOrganizations"`
	CompletedPrograms      []*orgs.Organization `bson: "completedPrograms" json:"completedPrograms"`
	InProcessPrograms      []*orgs.Organization `bson: "inProcessPrograms" json:"inProcessPrograms"`
	PathwayPrograms        []*orgs.Organization `bson: "pathwayPrograms" json: "pathwayPrograms"`
	AcademicPrograms       []*orgs.Organization `bson: "academicPrograms" json: "academicPrograms"`
	Notes                  []*Note              `bson: "notes" json:"notes"`
	Links                  []*Link              `bson: "links" json: "links"`
	Planners               []*Planner           `bson: "planners" json: "planners"`
	OrgsContacted          []*orgs.Org          `bson: "orgsContacted" json: "orgsContacted"`
}

// Credentials represents user sign-in credentials (student)
type Credentials struct {
	Email    string `bson: "email" json:"email"`
	Password string `bson: "password" json:"password"`
}

// NewUser represents a new user signing up for an account (student)
type NewUser struct {
	ID           primitive.ObjectID `"bson:_id"`
	Email        string             `bson: "email" json:"email"`
	PassHash     []byte             `bson:"password" json:"password"` //never JSON encoded/decoded
	PasswordConf []byte             `bson: "passwordConf" json:"passwordConf"`
	FirstName    string             `bson: "firstName" json:"firstName"`
	UserName     string             `bson: "userName" json:"userName"`
}

// Updates represents updates allowed to be edited by user (student)
type Updates struct {
	FirstName              string               `bson: "firstName" json:"firstName"`
	UserName               string               `bson: "userName" json:"userName"`
	State                  string               `bson: "state" json:"state"`
	FavoritedOrganizations []*orgs.Organization `bson: "favoritedOrganizations" json:"favoritedOrganizations"`
	CompletedPrograms      []*orgs.Organization `bson: "completedPrograms" json:"completedPrograms"`
	InProcessPrograms      []*orgs.Organization `bson: "inProcessPrograms" json:"inProcessPrograms"`
	UserPathwayPrograms    []*orgs.Organization `bson: "userPathwayPrograms" json: "userPathwayPrograms" `
	UserAcademicPrograms   []*orgs.Organization `bson: "userAcademicPrograms" json: "userAcademicPrograms"`
	UserNotes              []*Note              `bson: "userNotes" json:"userNotes"`
	UserLinks              []*Link              `bson: "userLinks" json: "userLinks"`
	UserPlanners           []*Planner           `bson: "userPlanners" json: "userPlanners"`
}

// Orgs represents the users' (student) organizations
// type Org struct {
// 	OrgID             primitive.ObjectID `bson: "_id" `
// 	OrgTitle          string             `bson: "orgTitle" json:"orgTitle"`
// 	OrgLocation       string             `bson: "orgLocation" json:"orgLocation"`
// 	OrgType           string             `bson: "orgType" json:"orgType"`
// 	OrgCompleting     bool               `bson: "orgCompleting" json:"orgCompleting"`
// 	OrgCompleted      bool               `bson: "orgCompleted" json:"orgCompleted"`
// 	IsPathwayProgram  bool               `bson: "isPathwayProgram" json: "isPathwayProgram"`
// 	IsAcademicProgram bool               `bson: "isAcademicProgram" json: "isAcademicProgram" `
// 	Contacted         bool               `bson: "contacted" json: "contacted"`
// 	OrgDescription    string             `bson: "orgDescription" json: "orgDescription"`
// }

// Notes represents users' (student) notes about programs
type Note struct {
	UserID          primitive.ObjectID `bson:"_id, omitempty"`
	NoteID          primitive.ObjectID `bson: "_id"`
	NoteContent     string             `bson: "noteContent" json:noteContent`
	OrgID           primitive.ObjectID `bson: "_id"`
	CreatedAt       string             `bson: "createdAt" json:createdAt`
	UpdatedAt       string             `bson: "updatedAt" json:updatedAt`
	Reviewed        bool               `bson: "reviewed" json:reviewed`
	NoteDescription string             `bson: "noteDescription" json: "noteDescription"`
}

type Link struct {
	LinkID          primitive.ObjectID `bson: "_id"`
	LinkDescription string             `bson: "linkDescription" json: "linkDescription" `
	Favorited       bool               `bson: "favorited" json: "favorited"`
	UserID          primitive.ObjectID `bson:"_id, omitempty"`
	PlannerID       primitive.ObjectID `bson: "_id"`
	OrgID           primitive.ObjectID `bson: "_id, omitempty"`
}

type Planner struct {
	PlannerID          primitive.ObjectID   `bson: "_id"`
	IsMonthlyPlanner   bool                 `bson: "isMonthlyPlanner" json: "isMonthlyPlanner"`
	IsYearlyPlanner    bool                 `bson: "isYearlyPlanner" json: "isYearlyPlanner"`
	IsAcademicPlanner  bool                 `bson: "isAcademicPlanner" "isAcademicPlanner"`
	NotesIDS           []*Note              `bson: "notesIDS" json: "notesIDS"`
	OrgIDS             []*orgs.Organization `bson: "orgIDS" json: "orgIDS"`
	UserID             primitive.ObjectID   `bson: "_id"`
	IsFallPlanner      bool                 `bson: "isFallPlanner" json: "isFallPlanner"`
	IsWinterPlanner    bool                 `bson: "isWinterPlanner" json: "isWinterPlanner"`
	IsSpringPlanner    bool                 `bson: "isSpringPlanner" json: "isSpringPlanner"`
	IsSummerPlanner    bool                 `bson: "isSummerPlanner" json: "isSummerPlanner"`
	LinkIDS            []*Link              `bson: "linkIDS" json: "linkIDS"`
	DateCreated        string               `bson: "dateCreated" json: "dateCreated"`
	PlannerDescription string               `bson: "plannerDescription" json: "plannerDescription"`
}

// UserOrgs represents a program administrator's organizations
// needed??
type UserOrgs struct {
	ID        int64                `json:"id"`
	Email     string               `json:"-"` //never JSON encoded/decoded
	FirstName string               `json:"firstName"`
	LastName  string               `json:"lastName"`
	Orgs      []*orgs.Organization `json:"orgs"`
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
