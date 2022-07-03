package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

func Run() {
	log.Println("API service started")
	setupRoutesAPI()
}

func setupRoutesAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", healthCheckAPI).Methods("GET")
	router.HandleFunc("/api/users", allUsersHandlerAPI).Methods("GET")
	router.HandleFunc("/api/user/{name}/{email}", newUserHandlerAPI).Methods("POST")
	router.HandleFunc("/api/user/{name}", deleteUserHandlerAPI).Methods("DELETE")
	router.HandleFunc("/api/user/{name}/{email}", updateUserHandlerAPI).Methods("PUT")

	port := os.Getenv("API_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}

func healthCheckAPI(rw http.ResponseWriter, req *http.Request) {
	log.Println("healthCheckAPI invoked")
	url := "http://db:8082/db/health"

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseMessage string
	if err = json.Unmarshal(bodyBytes, &responseMessage); err != nil {
		log.Println(err)
	}

	fmt.Fprintf(rw, fmt.Sprintf("%s. API service is working too", responseMessage))
}

func allUsersHandlerAPI(rw http.ResponseWriter, req *http.Request) {
	log.Println("allUsersHandlerAPI invoked")
	url := "http://db:8082/db/users"

	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseBody []userResponse
	if err = json.Unmarshal(bodyBytes, &responseBody); err != nil {
		log.Println(err)
	}

	log.Println(fmt.Sprintf("%d users fetched", len(responseBody)))
	json.NewEncoder(rw).Encode(responseBody)
}

func newUserHandlerAPI(rw http.ResponseWriter, req *http.Request) {
	log.Println("newUserHandlerAPI handler invoked")
	vars := mux.Vars(req)

	log.Println(vars)

	url := fmt.Sprintf("http://db:8082/db/user/%s/%s", vars["name"], vars["email"])
	response, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Println("1:", err)
	}

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println("2:", err)
	}

	var responseMessage string
	if err = json.Unmarshal(bodyBytes, &responseMessage); err != nil {
		log.Println("3:", err)
	}

	log.Println(responseMessage)
	fmt.Fprintf(rw, responseMessage)
}

func deleteUserHandlerAPI(rw http.ResponseWriter, req *http.Request) {
	log.Println("deleteUserHandlerAPI handler invoked")
	client := http.Client{}
	vars := mux.Vars(req)

	url := fmt.Sprintf("http://db:8082/db/user/%s", vars["name"])
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Println(err)
	}

	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		log.Println(err)
	}

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseMessage string
	if err = json.Unmarshal(bodyBytes, &responseMessage); err != nil {
		log.Println(err)
	}

	log.Println(responseMessage)
	fmt.Fprintf(rw, responseMessage)
}

func updateUserHandlerAPI(rw http.ResponseWriter, req *http.Request) {
	log.Println("updateUserHandlerAPI handler invoked")
	log.Println("deleteUserHandlerAPI handler invoked")
	client := http.Client{}
	vars := mux.Vars(req)

	url := fmt.Sprintf("http://db:8082/db/user/%s/%s", vars["name"], vars["email"])
	request, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		log.Println(err)
	}

	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		log.Println(err)
	}

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	var responseMessage string
	if err = json.Unmarshal(bodyBytes, &responseMessage); err != nil {
		log.Println(err)
	}

	log.Println(responseMessage)
	fmt.Fprintf(rw, responseMessage)
}
