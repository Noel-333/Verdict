package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

type SHA256Hasher struct{}

func (SHA256Hasher) Hash(File *os.File) string {
	defer File.Close()

	Hash := sha256.New()

	io.Copy(Hash, File)

	return hex.EncodeToString(Hash.Sum(nil))
}
