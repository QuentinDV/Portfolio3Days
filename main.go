package main

import (
	"Portfolio/assets/functions"
	"fmt"
	"net/http"
)

func main() {
	// Creating the database
	functions.ConnectAccountDB("db/database.db")
	functions.ConnectPortfolioDB("db/database.db")

	functions.InsertAccount("", "admin123", "Admin", "", true)
	Admin, _ := functions.GetAccountByEmail("")
	functions.InsertPortfolio(Admin.Id, Admin.Name, Admin.FamilyName, Admin.Email, "07 00 00 00 00", "linkedin.com", "github.com")

	functions.InsertAccount("quentin.dassivignon@ynov.com", "Quentin123", "Quentin", "Dassi Vignon", false)
	Account, _ := functions.GetAccountByEmail("quentin.dassivignon@ynov.com")
	functions.InsertPortfolio(Account.Id, Account.Name, Account.FamilyName, Account.Email, "06 00 00 00 00", "linkedin.com", "github.com")

	Portfolio, err := functions.GetPortfolio(Account.Id)
	if err != nil {
		fmt.Println("Error retrieving portfolio:", err)
	}

	// Create Project
	functions.InsertProjet(Portfolio.Id, "Hangman Web", "Un jeu de pendu en ligne avec différentes difficultés. Explore une interface intuitive pour défier tes connaissances en vocabulaire tout en t'amusant. La difficulté évolue selon le nombre de mots incorrects.", "HTML, CSS, JavaScript", "https://www.trainerbubble.com/wp-content/uploads/2015/09/Hangman_web-1024x682.jpg")
	functions.InsertProjet(Portfolio.Id, "Forum", "Un forum de discussion pour développeurs où partager des idées, résoudre des problèmes et discuter des dernières tendances en programmation. La communauté est active et propose des échanges variés.", "PHP, MySQL, HTML, CSS", "https://thumbs.dreamstime.com/b/forum-discutant-des-gens-23825945.jpg")
	functions.InsertProjet(Portfolio.Id, "Hangman", "Version classique du jeu de pendu. Le joueur doit deviner un mot en proposant des lettres. Si la lettre est incorrecte, une partie du dessin du pendu apparaît. Le but est de deviner le mot avant que le pendu soit complété.", "Python", "https://m.media-amazon.com/images/I/71RM2nXWstL.png")

	// Pages Handlers
	http.HandleFunc("/admin", functions.AdminPage)

	http.HandleFunc("/portfolio/", functions.PortfolioPageHome)
	http.HandleFunc("/project/", functions.PortfolioPageProject)
	http.HandleFunc("/delete/", functions.DeleteAccountForm)

	// Static Handlers
	http.Handle("/assets/css/", http.StripPrefix("/assets/css/", http.FileServer(http.Dir("./assets/css"))))
	http.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("./assets/img"))))

	// Creating the server
	fmt.Println("http://localhost:8080")
	fmt.Println("http://localhost:8080/admin")
	fmt.Println("http://localhost:8080/portfolio/1")
	fmt.Println("http://localhost:8080/project/1")
	http.ListenAndServe(":8080", nil)
}
