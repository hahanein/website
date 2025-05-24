package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello World"); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Println("listening on :8767")
	log.Fatal(http.ListenAndServe(":8767", nil))
}
