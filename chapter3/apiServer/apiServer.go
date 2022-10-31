package main

import (
	"go-implement-your-object-storage/chapter5/apiServer/heartbeat"
	"go-implement-your-object-storage/chapter5/apiServer/locate"
	"go-implement-your-object-storage/chapter5/apiServer/objects"
	"go-implement-your-object-storage/chapter5/apiServer/versions"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
