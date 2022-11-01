package objects

import (
	"go-implement-your-object-storage/chapter4/dataServer/locate"
	"go-implement-your-object-storage/src/lib/utils"
	"log"
	"net/url"
	"os"
)

func getFile(hash string) string {
	file := os.Getenv("STORAGE_ROOT") + "/objects/" + hash
	f, _ := os.Open(file)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}
