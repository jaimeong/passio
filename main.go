package main

import (
	"app/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "passio"
)

var users []models.User

func getUsers(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	//encode users into json
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)

	// loop thhru users, find with id
	// for _ + for item in users
	for _, item := range users {
		if item.Username == params["username"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&models.User{})

}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	insertDynStmt := `INSERT INTO users(username, passwordhash, firstname, middlename, lastname, email, phone) 
	values($1, $2, $3, $4, $5, $6, $7)`
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
	_, e := db.Exec(updateStmt, user.Password, user.Firstname, user.Middlename, user.Lastname, user.Email, user.Phone, params["username"])
	CheckError(e)

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)

	for index, item := range users {
		if item.Username == params["username"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
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
