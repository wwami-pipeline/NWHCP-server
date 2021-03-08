package users

import (
	"database/sql"
	"time"
)

type Database struct {
	DB *sql.DB
}

func GetNewStore(db *sql.DB) *Database {
	return &Database{db}
}

func (db *Database) GetByID(id int64) (*User, error) {
	row := db.DB.QueryRow("SELECT * FROM user WHERE user_id = ?", id)
	user := User{}
	if err := row.Scan(&user.ID, &user.Email, &user.PassHash, &user.FirstName, &user.LastName); err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

func (db *Database) GetByEmail(email string) (*User, error) {
	row := db.DB.QueryRow("SELECT * FROM user WHERE email = ?", email)
	user := User{}
	if err := row.Scan(&user.ID, &user.Email, &user.PassHash, &user.FirstName, &user.LastName); err != nil {
		return nil, ErrUserNotFound
	}
	return &user, nil
}

// Insert blah should I store as string or date?
func (db *Database) Insert(user *User) (*User, error) {
	insq := "INSERT INTO user(email, passhash, firstname, lastname, dob) VALUES (?,?,?,?,?)"
	birthDate, _ := time.Parse("2006-01-02", user.BirthDate)
	res, err := db.DB.Exec(insq, user.Email, user.PassHash, user.FirstName, user.LastName, birthDate)
	if err != nil {
		return nil, err
	}
	id, err2 := res.LastInsertId()
	if err2 != nil {
		return nil, err2
	}
	user.ID = id
	return user, nil
}

func (db *Database) Update(id int64, updates *Updates) (*User, error) {
	insq := "UPDATE user SET firstname = ?, lastname = ? WHERE user_id = ?"
	_, err := db.DB.Exec(insq, updates.FirstName, updates.LastName, id)
	if err != nil {
		return nil, ErrUserNotFound
	}
	user, err2 := db.GetByID(id)
	if err2 != nil {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (db *Database) Delete(id int64) error {
	insq := "DELETE FROM user WHERE user_id = ?"
	_, err := db.DB.Exec(insq, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) TrackLogin(id int64, ip string, time time.Time) error {
	query := "INSERT INTO SignIns(user_id, IPAddress, SignInDate) VALUES (?,?,?)"
	_, err := db.DB.Exec(query, id, ip, time)
	if err != nil {
		return err
	}
	return nil
}

<<<<<<< HEAD
// GetOrgs for SQL gets the Specific User's Organizations that they saved
// func (db *Database) GetOrgs(id int64) ([]*Orgs, error) {
func (db *Database) GetOrgs(id int64) (*UserOrgs, error) {
	// ("SELECT OrgID, OrgTitle FROM UserOrg UO JOIN Organization O on UO.OrgID = O.OrgID WHERE UserID = ?", userID)
	query := "SELECT OrgID, OrgTitle FROM UserOrg UO JOIN Organization O on UO.OrgID = O.OrgID WHERE UserID = ?"
=======
// GetOrgs for SQL
func (db *Database) GetOrgs(id int64) (*userOrgs, error) {
	// ("SELECT OrgID, OrgTitle FROM UserOrg UO JOIN Organization O on UO.OrgID = O.OrgID WHERE user_id = ?", user_id)
	query := "SELECT org_id, org_title FROM user_org UO JOIN organization O on UO.org_id = O.org_id WHERE user_id = ?"
>>>>>>> 3c658731ba38525d3e18e8a38d76991b66adf5c1
	orgs, error := db.DB.Query(query, id)
	if error != nil {
		return nil, error
	}
	// userOrgs[] =
	usr, err := db.GetByID(id)
	if err != nil {
		return nil, err
	}

	userOrgs := []*Orgs{}
	ret := &UserOrgs{}
	ret.ID = usr.ID
	ret.Email = usr.Email
	ret.FirstName = usr.FirstName
	ret.LastName = usr.LastName

	for orgs.Next() {
		temp := &Orgs{}
		if errRow := orgs.Scan(&temp.OrgID, &temp.OrgTitle); errRow != nil {
			// http.Error(w, "Database error", http.StatusInternalServerError)
			// return org{}, errors.New("Database error")
			return nil, errRow
		}
		userOrgs = append(userOrgs, temp)
	}

	ret.Orgs = userOrgs

	return ret, nil
	// return userOrgs, nil
}
