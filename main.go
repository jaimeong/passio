package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
	io.WriteString(w, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	http.HandleFunc("/", homePage)

	fmt.Println("Listening on port 8001...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
