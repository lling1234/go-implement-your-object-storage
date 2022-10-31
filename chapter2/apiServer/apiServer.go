package main

import (
	"go-implement-your-object-storage/chapter2/apiServer/heartbeat"
	"go-implement-your-object-storage/chapter2/apiServer/locate"
	"go-implement-your-object-storage/chapter2/apiServer/objects"
	"go-implement-your-object-storage/src/lib/common"
	"log"
	"net/http"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(common.HttpPort, nil))
}
