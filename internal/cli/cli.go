package cli

import (
	"io"
	"os"
)

func NewApp() *App {
	return &App{errStream: os.Stderr, outStream: os.Stdout}
}

type App struct {
	errStream io.Writer
	outStream io.Writer
}

func (a *App) Run(argv []string) int {
	return 0
}
