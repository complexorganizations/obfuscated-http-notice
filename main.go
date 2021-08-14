package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(handleRequest))
	http.ListenAndServe(":8080", nil)
}

func handleRequest(httpWriter http.ResponseWriter, r *http.Request) {
	// Set the header to status not found.
	httpWriter.WriteHeader(http.StatusNotFound)
	// Set the content type to application/json.
	httpWriter.Header().Set("Content-Type", "application/json")
	// Set the body to an error message.
	type errorMessage struct {
		Code    int
		Message string
	}
	errorMsg := errorMessage{
		Code:    http.StatusNotFound,
		Message: "Resource not found",
	}
	// Wrap the error in a error object.
	type jsonError struct {
		Error errorMessage
	}
	// The content of the error object.
	jsonReturn := jsonError{
		Error: errorMsg,
	}
	// Marshal the error message to JSON.
	errorJsonMessage, err := json.Marshal(jsonReturn)
	// Log the error if there is one.
	if err != nil {
		log.Println(err)
	}
	// Write the JSON error message.
	httpWriter.Write(errorJsonMessage)
}
