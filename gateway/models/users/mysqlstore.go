package users

// import (
// 	"database/sql"
// 	"time"
// )

// type Database struct {
// 	DB *sql.DB
// }

// func GetNewStore(db *sql.DB) *Database {
// 	return &Database{db}
// }

// func (db *Database) GetByID(id int64) (*User, error) {
// 	row := db.DB.QueryRow("SELECT * FROM Users WHERE UserID = ?", id)
// 	user := User{}
// 	if err := row.Scan(&user.ID, &user.Email, &user.PassHash, &user.FirstName, &user.LastName, &user.JoinDate); err != nil {
// 		return nil, ErrUserNotFound
// 	}
// 	return &user, nil
// }

// func (db *Database) GetByEmail(email string) (*User, error) {
// 	row := db.DB.QueryRow("SELECT * FROM Users WHERE Email = ?", email)
// 	user := User{}
// 	if err := row.Scan(&user.ID, &user.Email, &user.PassHash, &user.FirstName, &user.LastName, &user.JoinDate); err != nil {
// 		return nil, ErrUserNotFound
// 	}
// 	return &user, nil
// }

// // Insert blah should I store as string or date?
// func (db *Database) Insert(user *User) (*User, error) {
// 	insq := "INSERT INTO Users(Email, PassHash, FirstName, LastName, JoinDate) VALUES (?,?,?,?,?)"
// 	// birthDate, _ := time.Parse("2006-01-02", user.BirthDate)
// 	res, err := db.DB.Exec(insq, user.Email, user.PassHash, user.FirstName, user.LastName, user.JoinDate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	id, err2 := res.LastInsertId()
// 	if err2 != nil {
// 		return nil, err2
// 	}
// 	user.ID = id
// 	return user, nil
// }

// func (db *Database) Update(id int64, updates *Updates) (*User, error) {
// 	insq := "UPDATE Users SET FirstName = ?, LastName = ? WHERE UserID = ?"
// 	_, err := db.DB.Exec(insq, updates.FirstName, updates.LastName, id)
// 	if err != nil {
// 		return nil, ErrUserNotFound
// 	}
// 	user, err2 := db.GetByID(id)
// 	if err2 != nil {
// 		return nil, ErrUserNotFound
// 	}
// 	return user, nil
// }

// func (db *Database) Delete(id int64) error {
// 	insq := "DELETE FROM Users WHERE UserID = ?"
// 	_, err := db.DB.Exec(insq, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (db *Database) TrackLogin(id int64, ip string, time time.Time) error {
// 	query := "INSERT INTO SignIns(UserID, SignInDate, IPAddress) VALUES (?,?,?)"
// 	_, err := db.DB.Exec(query, id, time, ip)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // GetOrgs for SQL
// func (db *Database) GetOrgs(id int64) (*UserOrgs, error) {
// 	// ("SELECT OrgID, OrgTitle FROM UserOrg UO JOIN Organization O on UO.OrgID = O.OrgID WHERE user_id = ?", user_id)
// 	query := "SELECT org_id, org_title FROM user_org UO JOIN organization O on UO.org_id = O.org_id WHERE user_id = ?"
// 	orgs, error := db.DB.Query(query, id)
// 	if error != nil {
// 		return nil, error
// 	}
// 	// userOrgs[] =
// 	usr, err := db.GetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	userOrgs := []*Orgs{}
// 	ret := &UserOrgs{}
// 	ret.ID = usr.ID
// 	ret.Email = usr.Email
// 	ret.FirstName = usr.FirstName
// 	ret.LastName = usr.LastName

// 	for orgs.Next() {
// 		temp := &Orgs{}
// 		if errRow := orgs.Scan(&temp.OrgID, &temp.OrgTitle); errRow != nil {
// 			// http.Error(w, "Database error", http.StatusInternalServerError)
// 			// return org{}, errors.New("Database error")
// 			return nil, errRow
// 		}
// 		userOrgs = append(userOrgs, temp)
// 	}

// 	ret.Orgs = userOrgs

// 	return ret, nil
// 	// return userOrgs, nil
// }
