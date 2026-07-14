package fs

import (
	"errors"
	"os"
)

type FS interface {
	Ensure(FilePath string) error
	Create(FilePath string) error
	Open(FilePath string) (*os.File, error)
}

type OsFS struct{}

var (
	ErrFileAlreadyExists = errors.New("file already exists")
)

func (OsFS) Ensure(FilePath string) error {
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		if os.IsExist(err) {
			return ErrFileAlreadyExists
		}
		return err
	}
	return file.Close()
}

func (OsFS) Create(FilePath string) error {
	file, err := os.Create(FilePath)
	if err != nil {
		return err
	}
	return file.Close()
}

func (OsFS) Open(FilePath string) (*os.File, error) {
	file, err := os.OpenFile(FilePath, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
