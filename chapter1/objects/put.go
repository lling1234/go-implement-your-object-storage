package objects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)
var FileDIr string = "/home/ling/go/src/go-implement-your-object-storage/c1"
func put(w http.ResponseWriter, r *http.Request) {
	log.Println("r.Body",r.Body)
	log.Println("w",w)
	f, e := os.Create(FileDIr + "/objects/" +
		strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}
