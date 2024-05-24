package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Call the function
	Hello()

	// Create a new instance of the http server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a map to store the response
		response := map[string]interface{}{
			"message": "Hello, World!",
		}

		// Convert the map to json
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the content type to json
		w.Header().Set("Content-Type", "application/json")

		// Write the response
		_, err = w.Write(responseJSON)
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println(fmt.Sprintf("Error starting server: %s", err.Error()))
	}
}

func Hello() {
	println("Hello, World!")
}
