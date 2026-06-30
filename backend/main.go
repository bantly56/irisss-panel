package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Println("Irisss Panel started on :8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}

}
