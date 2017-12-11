package main

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tom1193/language-api/nlp"
	"log"
	//"github.com/tom1193/language-api/proto"
	"github.com/rs/cors"
)

func GetEntityEndpoint(w http.ResponseWriter, req *http.Request) {
	params := req.URL.Query()
	if params["text"] != nil {
		res, err := nlp.AnalyzeEntitySentiment(params["text"][0])
		if err != nil {
			log.Fatalf("Failed to analyze text: %v", err)
		}
		result := nlp.GenerateEntity(res)
		json.NewEncoder(w).Encode(result)
	} else {
		log.Fatalf("Failed to parse query params")
	}
}

func main () {
	router := mux.NewRouter()
	handler := cors.Default().Handler(router)
	router.HandleFunc("/entity", GetEntityEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
