package main

import (
	"app/models"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
	host     = "172.17.0.1"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "passio"
)

var users []models.User

func getUsers(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(`SELECT username, passwordhash, firstname, middlename, lastname, email, phone FROM users`)
	CheckError(err)

	var users []models.User

	defer rows.Close()
	for rows.Next() {
		var username string
		var firstname string
		var middlename string
		var lastname string
		var email string
		var phone string
		var passwordhash string

		err = rows.Scan(&username, &passwordhash, &firstname, &middlename, &lastname, &email, &phone)
		CheckError(err)

		users = append(users, models.User{
			Username:   username,
			Password:   passwordhash,
			Firstname:  firstname,
			Middlename: middlename,
			Lastname:   lastname,
			Email:      email,
			Phone:      phone,
		})
	}

	CheckError(err)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	getStmt := `SELECT username, passwordhash, firstname, middlename, lastname, email, phone FROM users where username = $1`
	rows, err := db.Query(getStmt, params["username"])
	var users []models.User

	defer rows.Close()
	for rows.Next() {
		var username string
		var firstname string
		var middlename string
		var lastname string
		var email string
		var phone string
		var passwordhash string

		err = rows.Scan(&username, &passwordhash, &firstname, &middlename, &lastname, &email, &phone)
		CheckError(err)

		users = append(users, models.User{
			Username:   username,
			Password:   passwordhash,
			Firstname:  firstname,
			Middlename: middlename,
			Lastname:   lastname,
			Email:      email,
			Phone:      phone,
		})
	}

	CheckError(err)
	json.NewEncoder(w).Encode(users)
}

func hashPass(password string) string {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.
		GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)

	if err != nil {
		log.Println(err)
	}

	// Convert the hashed password to a base64 encoded string
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	insertDynStmt := `INSERT INTO users(username, passwordhash, firstname, middlename, lastname, email, phone) 
	values($1, $2, $3, $4, $5, $6, $7)`
	user.Password = hashPass(user.Password)
	_, err = db.Exec(insertDynStmt, user.Username, user.Password, user.Firstname, user.Middlename, user.Lastname, user.Email, user.Phone)
	CheckError(err)
	json.NewEncoder(w).Encode(user)

}

// // combintation of delete + create
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	params := mux.Vars(r)
	updateStmt := `update users set passwordhash = $1, firstname = $2, middlename = $3, lastname = $4, email = $5, phone = $6

	where "username"=$7`

	user.Password = hashPass(user.Password)
	_, e := db.Exec(updateStmt, user.Password, user.Firstname, user.Middlename, user.Lastname, user.Email, user.Phone, params["username"])
	CheckError(e)

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)

	deleteStmt := `delete from users where username = $1`
	_, e := db.Exec(deleteStmt, params["username"])
	CheckError(e)

	fmt.Fprintf(w, "Account deleted \n")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

var db *sql.DB
var err error

func main() {
	router := mux.NewRouter()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/user/{username}", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")
	router.HandleFunc("/api/user/{username}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{username}", deleteUser).Methods("DELETE")

	fmt.Println("Listening on port 8001...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
