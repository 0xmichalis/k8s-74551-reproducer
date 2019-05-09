package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		if r.Method == http.MethodPost {
			fmt.Printf("Content length of POST req: %d\n", r.ContentLength)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
