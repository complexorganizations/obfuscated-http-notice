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
	httpWriter.WriteHeader(http.StatusNotFound)
	httpWriter.Header().Set("Content-Type", "application/json")
	type errorMessage struct {
		Message string
	}
	errorMsg := errorMessage{
		Message: "Resource Not Found",
	}
	errorJsonMessage, err := json.Marshal(errorMsg)
	if err != nil {
		log.Println("error:", err)
	}
	httpWriter.Write(errorJsonMessage)
}
