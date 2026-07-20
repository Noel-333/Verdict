package hasher

import (
	"encoding/hex"
	"io"
	"os"

	"github.com/cespare/xxhash/v2"
)

type XxhashHasher struct{}

func (XxhashHasher) Hash(File *os.File) string {
	defer File.Close()

	Hash := xxhash.New()

	io.Copy(Hash, File)

	return hex.EncodeToString(Hash.Sum(nil))
}
