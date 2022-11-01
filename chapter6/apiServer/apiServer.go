package main

import (
	"go-implement-your-object-storage/chapter6/apiServer/heartbeat"
	"go-implement-your-object-storage/chapter6/apiServer/locate"
	"go-implement-your-object-storage/chapter6/apiServer/objects"
	"go-implement-your-object-storage/chapter6/apiServer/temp"
	"go-implement-your-object-storage/chapter6/apiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
