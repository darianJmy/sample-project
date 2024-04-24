package main

import (
	"os"
	"sample-project/cmd/app"
)

func main() {
	cmd := app.NewSampleServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
