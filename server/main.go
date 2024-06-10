package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from backend server on port %s!", port)
	})

	fmt.Printf("Starting backend server on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
