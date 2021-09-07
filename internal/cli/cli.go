package cli

import (
	"flag"
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func NewApp() *App {
	return &App{errStream: os.Stderr, outStream: os.Stdout}
}

type App struct {
	errStream io.Writer
	outStream io.Writer
}

func (a *App) Run(argv []string) int {
	fs := flag.NewFlagSet(argv[0], flag.ContinueOnError)
	fs.SetOutput(a.errStream)
	err := fs.Parse(argv[1:])
	if err == flag.ErrHelp {
		return 0
	}
	if err != nil {
		log.Warn().Err(err).Send()
		return 1
	}

	return 0
}
