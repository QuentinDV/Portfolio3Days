package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {

}

func AdminPage(w http.ResponseWriter, r *http.Request) {
	// Creating the database
	err := ConnectAccountDB("db/database.db")
	if err != nil {
		fmt.Println("Error connecting to Account DB:", err)
		return
	}

	AllAccounts, err := GetAllAccounts()
	if err != nil {
		fmt.Println("Error retrieving accounts:", err)
		return
	}

	// Serve the admin page
	tmpl := template.Must(template.ParseFiles("assets/html/admin.html"))
	tmpl.Execute(w, AllAccounts[1:])
}

func PortfolioPageHome(w http.ResponseWriter, r *http.Request) {
	// Extract the Portofolio ID from the URL
	PortofolioID := r.URL.Path[len("/portfolio/"):]

	Portofolio, err := GetPortfolio(PortofolioID)
	if err != nil {
		fmt.Println("Error retrieving Portofolio:", err)
		return
	}

	// Execute the Portfolio template with the struct
	tmpl := template.Must(template.ParseFiles("assets/html/home.html"))
	tmpl.Execute(w, Portofolio)
}

func PortfolioPageProject(w http.ResponseWriter, r *http.Request) {
	// Extract the Project ID from the URL
	ProjectID := r.URL.Path[len("/project/"):]

	Project, err := GetProject(ProjectID)
	if err != nil {
		fmt.Println("Error retrieving Project:", err)
		return
	}

	fmt.Println("Project:", Project)

	// Execute the Project template with the struct
	tmpl := template.Must(template.ParseFiles("assets/html/project.html"))
	tmpl.Execute(w, Project)
}

func DeleteAccountForm(w http.ResponseWriter, r *http.Request) {
	// Extract the Account ID from the URL
	AccountID := r.URL.Path[len("/delete/"):]

	err := DeleteAccount(AccountID)
	if err != nil {
		fmt.Println("Error deleting account:", err)
		return
	}

	// Redirect to the admin page
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
