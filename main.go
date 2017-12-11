package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tom1193/language-api/nlp"
	"log"
)

func GetEntityEndpoint(w http.ResponseWriter, req *http.Request) {
	//params := mux.Vars(req)
	res, err := nlp.AnalyzeEntitySentiment("Michelangelo Caravaggio, Italian painter, is known for 'The Calling of Saint Matthew'.")
	if err != nil {
		log.Fatalf("Failed to analyze text: %v", err)
	}
	result := nlp.GenerateEntity(res)
	json.NewEncoder(w).Encode(result)
}

func main () {
	router := mux.NewRouter()
	router.HandleFunc("/entity", GetEntityEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
