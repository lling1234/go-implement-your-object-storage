package main

import (
	"go-implement-your-object-storage/chapter6/dataServer/heartbeat"
	"go-implement-your-object-storage/chapter6/dataServer/locate"
	"go-implement-your-object-storage/chapter6/dataServer/objects"
	"go-implement-your-object-storage/chapter6/dataServer/temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
