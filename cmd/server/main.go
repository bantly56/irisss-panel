package main

import (
	"log"
	"net/http"

	"github.com/bantly56/irisss-panel/internal/api"
)

func main() {

	router := api.NewRouter()

	log.Println("Irisss Panel started on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
