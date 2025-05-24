package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	log.Println("listening on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}
