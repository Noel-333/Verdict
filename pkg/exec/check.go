package exec

import (
	"encoding/json"
	"io"
	"slices"

	"github.com/fatih/color"
	"main.go/internal/version"
	"main.go/pkg/infra/hasher"
)

func CheckMeta() bool {
	file, err = f.Open("verdict.json")

	if err != nil {
		color.Red("verdict.json load failed: %v", err)
		return false
	}
	defer file.Close()

	data, err = io.ReadAll(file)
	if err != nil {
		color.Red("read verdict.json: %v", err)
		return false
	}

	err = json.Unmarshal(data, &manifest)
	if err != nil {
		color.Red("parse verdict.json failed: %v", err)
		return false
	}

	if manifest.Meta.Version != version.Version {
		color.Red("the configuration file version is %s which is not supported", manifest.Meta.Version)
		return false
	}

	if !slices.Contains(hasher.SupportedAlgo, manifest.Meta.Algorithm) {
		color.Red("algorithm %s not supported", manifest.Meta.Algorithm)
		return false
	}

	return true
}

func CheckData() {

	j := manifest.Data

	for i := range manifest.Data {

		h, err := hasher.New(manifest.Meta.Algorithm)

		if err != nil {
			color.Red("failed to initialize hasher for algorithm: %s", manifest.Meta.Algorithm)
		}

		k, err := f.Open(j[i].File)

		if err != nil {
			color.Red("file %s not exists :")
		} else {
			out := h.Hash(k)
			if out == j[i].Hash {
				color.Green("%s:%s", j[i].File, out)
			} else {
				color.Red("%s:%s", j[i].File, out)
			}

		}

	}
}
