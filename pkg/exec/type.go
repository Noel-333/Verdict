package exec

import (
	"os"

	"main.go/pkg/infra/fs"
)

var (
	err      error
	f        fs.FS = fs.OsFS{}
	file     *os.File
	data     []byte
	manifest Manifest
	Wd       string
)

type FileEntry struct {
	File string `json:"file"`
	Hash string `json:"hash"`
}

type Manifest struct {
	Meta struct {
		Version   string `json:"version"`
		Algorithm string `json:"algorithm"`
	} `json:"meta"`
	Data []FileEntry `json:"data"`
}

func init() {
	Wd, _ = os.Getwd()
}
