package exec

import (
	"encoding/json"
	"fmt"
	iofs "io/fs"
	"path/filepath"

	"github.com/fatih/color"
	"main.go/internal"
	"main.go/internal/version"
	"main.go/pkg/infra/fs"
	"main.go/pkg/infra/hasher"
)

func Snapshot() {

	if f.Ensure("verdict.json") == fs.ErrFileAlreadyExists {
		color.Yellow("verdict.json already exists. Continue? [Y/N]")

		var input string
		fmt.Scanln(&input)
		if input == "y" || input == "Y" {
			color.Green("")
		} else {
			color.Red("")
		}

	}

	var snapshot_path []string //snapshot处理的目录

	filepath.WalkDir(Wd, func(path string, d iofs.DirEntry, err error) error {
		if d.Name() == "vec.exe" || d.Name() == "verdict.json" {
			return nil
		}

		if !d.IsDir() {
			snapshot_path = append(snapshot_path, path)
			return nil
		}

		return nil
	})

	h, err := hasher.New(internal.SnapshotAValue)
	if err != nil {
		color.Red("failed to initialize hasher:%v", err)
		return
	}

	var Hash_result []FileEntry

	for i := range snapshot_path {
		var path string

		path = snapshot_path[i]

		file, err := f.Open(path)
		if err != nil {
			color.Red("failed to open file `%s` :%s", path, err)
			return
		}

		result := h.Hash(file)

		Hash_result = append(Hash_result, FileEntry{File: file.Name(), Hash: result})

	}

	m := Manifest{
		Meta: struct {
			Version   string "json:\"version\""
			Algorithm string "json:\"algorithm\""
		}{
			Version:   version.Version,
			Algorithm: internal.SnapshotAValue,
		},
		Data: Hash_result,
	}

	j, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		color.Red("failed to marshal:%v", err)
		return
	}

	err = f.Create("verdict.json")
	if err != nil {
		color.Red("failed to create verdict.json:%v", err)
		return
	}

	file, err := f.Open("verdict.json")
	if err != nil {
		color.Red("failed to open verdict.json:%v", err)
	}

	n, err := file.Write(j)
	if n == 0 || err != nil {
		color.Red("failed to write to verdict.json:%v", err)
	}

}
