package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var commands = []Command{Command{"tar", "an archiving utility"}}

type Command struct {
	Title       string
	Description string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/commands", getCommandsHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getCommandsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(commands)
	w.Write(j)
}
