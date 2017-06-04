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
	var err error;
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=flash");
	db.Exec("INSERT INTO main_schema.object_type VALUES ('card'), ('deck'), ('collection') ON CONFLICT DO NOTHING")
	db.Exec("INSERT INTO main_schema.card_action VALUES ('insert'), ('delete') ON CONFLICT DO NOTHING")
	db.Exec("INSERT INTO main_schema.deck_action VALUES ('rename'), ('insert_card'), ('remove_card') ON CONFLICT DO NOTHING")
	db.Exec("INSERT INTO main_schema.collection_action VALUES ('rename'), ('insert_deck'), ('remove_deck') ON CONFLICT DO NOTHING")
	
	if err != nil {
		log.Fatal(err);
	}

	hub := newHub();
	go hub.run();

	r := mux.NewRouter();
	r.Handle("/", http.FileServer(http.Dir("./views/")));
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r);
	})

	var port int = 3000;
	var listen_path string = fmt.Sprintf(":%d", port);
	err = http.ListenAndServe(listen_path, r);

	if err != nil {
		log.Fatal("ListenAndServe: ", err);
	}
}
