package main

import (
	"go-implement-your-object-storage/chapter2/dataServer/heartbeat"
	"go-implement-your-object-storage/chapter2/dataServer/locate"
	"go-implement-your-object-storage/chapter2/dataServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
