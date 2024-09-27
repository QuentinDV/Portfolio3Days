package functions

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Account struct represents a user account in the system
type Account struct {
	Id           string
	Email        string
	Password     string
	Name         string
	FamilyName   string
	CreationDate string
	IsAdmin      bool
}

func ConnectAccountDB(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
	id TEXT PRIMARY KEY,
	email TEXT UNIQUE NOT NULL,
	password TEXT NOT NULL,
	name TEXT NOT NULL,
	family_name TEXT NOT NULL,
	creation_date TEXT NOT NULL,
	is_admin BOOLEAN NOT NULL
	);`)
	if err != nil {
		return err
	}

	return nil
}

func InsertAccount(email, password, name, familyName string, isAdmin bool) error {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return err
	}

	lastID, err := GetLastId("accounts")
	if err != nil {
		return err
	}

	creationDate := time.Now().Format("02/01/2006")

	_, err = db.Exec(`INSERT INTO accounts (id, email, password, name, family_name, creation_date, is_admin)
        VALUES (?, ?, ?, ?, ?, ?, ?)`, lastID, email, password, name, familyName, creationDate, isAdmin)
	return err
}

func GetLastId(database string) (string, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return "", err
	}

	row := db.QueryRow("SELECT MAX(id) FROM " + database)
	var lastID sql.NullInt64
	err = row.Scan(&lastID)
	if err != nil {
		return "", err
	}

	if lastID.Valid {
		nextID := int(lastID.Int64) + 1
		return strconv.Itoa(nextID), nil
	}

	// If lastID is not valid, return "0"
	return "0", nil
}

func GetAccount(id string) (*Account, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`SELECT id, email, password, name, family_name, creation_date, is_admin
        FROM accounts WHERE id = ?`, id)

	account := &Account{}
	err = row.Scan(&account.Id, &account.Email, &account.Password, &account.Name, &account.FamilyName, &account.CreationDate, &account.IsAdmin)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func GetAccountByEmail(email string) (*Account, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return nil, err
	}

	row := db.QueryRow(`SELECT id, email, password, name, family_name, creation_date, is_admin
		FROM accounts WHERE email = ?`, email)

	account := &Account{}
	err = row.Scan(&account.Id, &account.Email, &account.Password, &account.Name, &account.FamilyName, &account.CreationDate, &account.IsAdmin)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func GetAllAccounts() ([]*Account, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT id, email, password, name, family_name, creation_date, is_admin FROM accounts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []*Account{}
	for rows.Next() {
		account := &Account{}
		err = rows.Scan(&account.Id, &account.Email, &account.Password, &account.Name, &account.FamilyName, &account.CreationDate, &account.IsAdmin)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func DeleteAccount(id string) error {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM accounts WHERE id = ?", id)
	return err
}
