package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", helloHandler).Methods("GET")
	fmt.Println("Server running at :8080")
	http.ListenAndServe(":8080", router)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Kubernetes")
	msg := map[string]string{
		"message": "Hello World",
	}
	json.NewEncoder(w).Encode(msg)
}
