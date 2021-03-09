package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./site"))
	http.Handle("/", fs)

	log.Println("forntends up")
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal(err)
	}
}
