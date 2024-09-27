package functions

import "database/sql"

type Portfolio struct {
	Id         string
	AccountId  string
	Name       string
	FamilyName string
	Email      string
	Phone      string
	Linkedin   string
	Github     string
	Projets    []Project
}

type Project struct {
	Id          string
	PortfolioId string
	Name        string
	Description string
	Technos     string
	Images      string
}

func ConnectPortfolioDB(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS portfolios (
	id TEXT PRIMARY KEY,
	account_id TEXT UNIQUE NOT NULL,
	name TEXT NOT NULL,
	family_name TEXT NOT NULL,
	email TEXT NOT NULL,
	phone TEXT NOT NULL,
	linkedin TEXT NOT NULL,
	github TEXT NOT NULL
	);`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS projets (
	id TEXT PRIMARY KEY,
	portfolio_id TEXT NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	technos TEXT NOT NULL,
	images TEXT NOT NULL,
	FOREIGN KEY (portfolio_id) REFERENCES portfolios(id)
	);`)
	if err != nil {
		return err
	}

	return nil
}

func InsertPortfolio(accountId, name, familyName, email, phone, linkedin, github string) error {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return err
	}

	lastID, err := GetLastId("portfolios")
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO portfolios (id, account_id, name, family_name, email, phone, linkedin, github)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, lastID, accountId, name, familyName, email, phone, linkedin, github)
	return err
}

func GetPortfolio(accountId string) (Portfolio, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return Portfolio{}, err
	}

	rows, err := db.Query(`SELECT * FROM portfolios WHERE account_id = ?`, accountId)
	if err != nil {
		return Portfolio{}, err
	}
	defer rows.Close()

	var portfolio Portfolio
	for rows.Next() {
		err = rows.Scan(&portfolio.Id, &portfolio.AccountId, &portfolio.Name, &portfolio.FamilyName, &portfolio.Email, &portfolio.Phone, &portfolio.Linkedin, &portfolio.Github)
		if err != nil {
			return Portfolio{}, err
		}
	}

	return portfolio, nil
}

func InsertProjet(portfolioId, name, description, technos, images string) error {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return err
	}

	lastID, err := GetLastId("projets")
	if err != nil {
		return err
	}

	_, err = db.Exec(`INSERT INTO projets (id, portfolio_id, name, description, technos, images)
		VALUES (?, ?, ?, ?, ?, ?)`, lastID, portfolioId, name, description, technos, images)
	return err
}

func GetProject(portfolioId string) ([]Project, error) {
	db, err := sql.Open("sqlite3", "db/database.db")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(`SELECT * FROM projets WHERE portfolio_id = ?`, portfolioId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projets := []Project{}
	for rows.Next() {
		var project Project
		err = rows.Scan(&project.Id, &project.PortfolioId, &project.Name, &project.Description, &project.Technos, &project.Images)
		if err != nil {
			return nil, err
		}
		projets = append(projets, project)
	}

	return projets, nil
}
