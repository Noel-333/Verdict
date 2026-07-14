package exec

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"main.go/pkg/infra/fs"
)

func Initialize() {
	err := f.Ensure("verdict.json")

	if err == nil {
		fmt.Println("file created successfully")
		color.Green("verdict initialized successfully")
		return
	}

	if errors.Is(err, fs.ErrFileAlreadyExists) {
		fmt.Print("verdict.json already exists. Overwrite? [y/N]")
		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "Y" {
			if createErr := f.Create("verdict.json"); createErr != nil {
				color.Red("file overwritten failed: %v", createErr)
			} else {
				color.Green("file overwritten successfully")
			}
		} else {
			color.Red("file not overwritten")
		}
		return
	}

	color.Red("Verdict initialization failed: %v", err)
}
