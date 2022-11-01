package objects

import (
	"fmt"
	"go-implement-your-object-storage/chapter2/apiServer/locate"
	"go-implement-your-object-storage/src/lib/objectstream"
	"io"
)

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}
	return objectstream.NewGetStream(server, object)
}
