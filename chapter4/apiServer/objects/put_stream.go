package objects

import (
	"fmt"
	"go-implement-your-object-storage/chapter4/apiServer/heartbeat"
	"go-implement-your-object-storage/src/lib/objectstream"
)

func putStream(hash string, size int64) (*objectstream.TempPutStream, error) {
	server := heartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any dataServer")
	}

	return objectstream.NewTempPutStream(server, hash, size)
}
