package objects

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"go-implement-your-object-storage/chapter8/dataServer/locate"

)

func del(w http.ResponseWriter, r *http.Request) {
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	files, _ := filepath.Glob(os.Getenv("STORAGE_ROOT") + "/objects/" + hash + ".*")
	if len(files) != 1 {
		return
	}
	locate.Del(hash)
	os.Rename(files[0], os.Getenv("STORAGE_ROOT")+"/garbage/"+filepath.Base(files[0]))
}
