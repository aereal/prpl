package main

import (
	"os"

	"github.com/aereal/prpl/internal/cli"
)

func main() {
	app := cli.NewApp()
	os.Exit(app.Run(os.Args))
}
