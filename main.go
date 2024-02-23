package main

import (
	"log"
	"net/http"
)

func myHandler(myWriter http.ResponseWriter, myRequest *http.Request) {
	myWriter.Header().Set("Content-Type", "application/json")
	switch myRequest.Method {
	case "GET":
		myWriter.WriteHeader(http.StatusOK)
		myWriter.Write([]byte(`{"message": "hello world"}`))
	case "POST":
		myWriter.WriteHeader(http.StatusCreated)
		myWriter.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		myWriter.WriteHeader(http.StatusAccepted)
		myWriter.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		myWriter.WriteHeader(http.StatusOK)
		myWriter.Write([]byte(`{"message": "delete called"}`))
	default:
		myWriter.WriteHeader(http.StatusNotFound)
		myWriter.Write([]byte(`{"message": "not found"}`))
	}
}
func main() {
	http.HandleFunc("/", myHandler)

	log.Println("Started server at localhost:8080")

	err := http.ListenAndServe(":8080", nil) // Passing nil since http.Handler has been added
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
