package orgs

import (
	"database/sql"
	"errors"
)

type Database struct {
	DB *sql.DB
}

var ErrOrgNotFound = errors.New("organization not found")

func GetNewStore(db *sql.DB) *Database {
	return &Database{db}
}

func (db *Database) GetByID(id int64) (*Organization, error) {
	row := db.DB.QueryRow("SELECT * FROM organization WHERE org_id = ?", id)
	org := Organization{}
	if err := row.Scan(&org.OrgId, &org.OrgTitle); err != nil {
		return nil, ErrOrgNotFound
	}
	return &org, nil
}

func (db *Database) GetByTitle(orgTitle string) (*Organization, error) {
	row := db.DB.QueryRow("SELECT * FROM organization WHERE org_title = ?", orgTitle)
	org := Organization{}
	if err := row.Scan(&org.OrgId, &org.OrgTitle); err != nil {
		return nil, ErrOrgNotFound
	}
	return &org, nil
}

func (db *Database) Insert(org *Organization) (*Organization, error) {
	insq := "INSERT INTO organization(org_title) VALUES (?)"
	res, err := db.DB.Exec(insq, org.OrgTitle)
	if err != nil {
		return nil, err
	}
	id, err2 := res.LastInsertId()
	if err2 != nil {
		return nil, err2
	}
	// consider changing in the int in struct to int64
	org.OrgId = int(id)
	return org, nil
}

func (db *Database) Update(id int64, newTitle string) (*Organization, error) {
	insq := "UPDATE organization SET org_title = ? WHERE org_id = ?"
	_, err := db.DB.Exec(insq, newTitle, id)
	if err != nil {
		return nil, ErrOrgNotFound
	}
	org, err2 := db.GetByID(id)
	if err2 != nil {
		return nil, ErrOrgNotFound
	}
	return org, nil
}

func (db *Database) Delete(id int64) error {
	insq := "DELETE FROM organization WHERE org_id = ?"
	_, err := db.DB.Exec(insq, id)
	if err != nil {
		return err
	}
	return nil
}