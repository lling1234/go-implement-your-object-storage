package main

import (
	"log"
	"net/http"

	"go-implement-your-object-storage/chapter1/objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe("2323", nil))
}
