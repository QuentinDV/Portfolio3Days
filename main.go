package main

import (
	"Portfolio/functions"
	"fmt"
	"net/http"
)

func main() {
	// Creating the database
	err := functions.ConnectAccountDB("db/database.db")
	if err != nil {
		fmt.Println("Error connecting to Account DB:", err)
		return
	}

	err = functions.ConnectPortfolioDB("db/database.db")
	if err != nil {
		fmt.Println("Error connecting to Portfolio DB:", err)
		return
	}

	err = functions.InsertAccount("", "admin123", "Admin", "", true)
	if err != nil {
		fmt.Println("Error inserting admin account:", err)
	}

	err = functions.InsertAccount("quentin.dassivignon@ynov.com", "Quentin123", "Quentin", "Dassi Vignon", false)
	if err != nil {
		fmt.Println("Error inserting Quentin account:", err)
	}

	Account, err := functions.GetAccountByEmail("quentin.dassivignon@ynov.com")
	if err != nil {
		fmt.Println("Error retrieving account:", err)
		return
	}
	if Account == nil {
		fmt.Println("Account not found")
		return
	}

	functions.InsertPortfolio(Account.Id, Account.Name, Account.FamilyName, Account.Email, "06 00 00 00 00", "linkedin.com", "github.com")
	if err != nil {
		fmt.Println("Error inserting portfolio:", err)
		return
	}

	Portfolio, err := functions.GetPortfolio(Account.Id)
	if err != nil {
		fmt.Println("Error retrieving portfolio:", err)
		return
	}

	fmt.Println("Portfolio:", Portfolio)
	fmt.Println("Portfolio.Id:", Portfolio.Id)

	// Create Project
	err = functions.InsertProjet(Portfolio.Id, "Tic Tac Toe", "Made a Tic Tac Toe", "Golang", "img/tictactoe.jpg")
	if err != nil {
		fmt.Println("Error inserting project:", err)
		return
	}

	// Creating the server
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
