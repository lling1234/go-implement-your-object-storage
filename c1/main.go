package main

import (
	"go-implement-your-object-storage/chapter1/objects"
	"log"
	"net/http"
)

func main() {
	log.Println("server listening 1212~~~~")

	http.HandleFunc("/objects/", objects.Handler)
	http.ListenAndServe(":1212", nil)
	log.Println("server listening 1212```````")
}
