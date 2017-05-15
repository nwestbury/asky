package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/lib/pq"
)

// Make db a global variable
var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=flash")
	
	if err != nil {
		log.Fatal(err)
	}
	
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	r.Handle("/login", LoginHandler).Methods("GET")
	r.Handle("/register", RegisterHandler).Methods("GET")
	r.Handle("/status", StatusHandler).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/fb", handleFacebookLogin).Methods("GET")

	var port int = 3000;
	var listen_path string = fmt.Sprintf(":%d", port);	
	http.ListenAndServe(listen_path, r)
}
