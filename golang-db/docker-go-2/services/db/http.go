package db

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Run() {
	log.Println("DB service started")
	initialMigration()
	setupRoutesDB()
}

func setupRoutesDB() {
	router := mux.NewRouter()

	router.HandleFunc("/db/health", healthCheckDB).Methods("GET")
	router.HandleFunc("/db/users", allUsersHandlerDB).Methods("GET")
	router.HandleFunc("/db/user/{name}/{email}", newUserHandlerDB).Methods("POST")
	router.HandleFunc("/db/user/{name}", deleteUserHandlerDB).Methods("DELETE")
	router.HandleFunc("/db/user/{name}/{email}", updateUserHandlerDB).Methods("PUT")

	port := os.Getenv("DB_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func healthCheckDB(rw http.ResponseWriter, req *http.Request) {
	log.Println("healthCheckDB invoked")

	responseMessage := "DB service is working"
	json.NewEncoder(rw).Encode(responseMessage)
}

func allUsersHandlerDB(rw http.ResponseWriter, req *http.Request) {
	log.Println("allUsersHandlerDB handler invoked")

	users := allUsersDB()
	json.NewEncoder(rw).Encode(users)
}

func newUserHandlerDB(rw http.ResponseWriter, req *http.Request) {
	log.Println("newUserHandlerDB invoked")

	vars := mux.Vars(req)
	responseMessage := newUserDB(vars["name"], vars["email"])

	json.NewEncoder(rw).Encode(responseMessage)
}

func deleteUserHandlerDB(rw http.ResponseWriter, req *http.Request) {
	log.Println("deleteUserHandlerDB invoked")

	vars := mux.Vars(req)
	responseMessage := deleteUserDB(vars["name"])

	json.NewEncoder(rw).Encode(responseMessage)
}

func updateUserHandlerDB(rw http.ResponseWriter, req *http.Request) {
	log.Println("updateUserHandlerDB invoked")

	vars := mux.Vars(req)
	responseMessage := updateUserDB(vars["name"], vars["email"])

	json.NewEncoder(rw).Encode(responseMessage)
}
