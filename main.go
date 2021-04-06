package main

import (
	"app/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

	users = append(users, user)
	json.NewEncoder(w).Encode(user)

}

// // combintation of delete + create
func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get params
	params := mux.Vars(r)

	for index, item := range users {
		if item.Username == params["username"] {
			// delete's slice
			users = append(users[:index], users[index+1:]...)

			// create's creation
			var user models.User
			_ = json.NewDecoder(r.Body).Decode(&user)

			user.Username = params["username"]
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
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

func main() {
	router := mux.NewRouter()

	users = append(users, models.User{
		Username:   "1",
		Password:   "$5F3fAz1",
		Firstname:  "John",
		Lastname:   "Smith",
		Middlename: "Kane",
		Email:      "test@gmail.com",
		Phone:      "123-456-7890",
	})

	users = append(users, models.User{
		Username:   "2",
		Password:   "$45zb34azza1",
		Firstname:  "Jane",
		Lastname:   "Doe",
		Middlename: "Zae",
		Email:      "test33@gmail.com",
		Phone:      "000-000-7890",
	})

	router.HandleFunc("/api/users", getUsers).Methods("GET")
	router.HandleFunc("/api/user/{username}", getUser).Methods("GET")
	router.HandleFunc("/api/user", createUser).Methods("POST")
	router.HandleFunc("/api/user/{username}", updateUser).Methods("PUT")
	router.HandleFunc("/api/user/{username}", deleteUser).Methods("DELETE")

	fmt.Println("Listening on port 8001...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
