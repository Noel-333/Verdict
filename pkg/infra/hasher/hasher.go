package hasher

import (
	"fmt"
	"os"
)

var (
	SupportedAlgo = []string{"sha256"}
)

type Hasher interface {
	Hash(File *os.File) string
}

func New(hasherType string) (Hasher, error) {
	switch hasherType {
	case "sha256":
		return &SHA256Hasher{}, nil
	default:
		if hasherType == "" {
			return &SHA256Hasher{}, nil
		}

		err := fmt.Errorf("no such algorithm: %s", hasherType)
		return nil, err
	}
}
