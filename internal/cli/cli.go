package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"

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

	ctx := context.Background()
	if err := a.run(ctx, fs.Args()); err != nil {
		log.Error().Err(err).Send()
		return 1
	}

	return 0
}

func (a *App) run(ctx context.Context, args []string) error {
	log.Debug().Strs("args", args).Send()
	if len(args) == 0 {
		return fmt.Errorf("the command must be given")
	}

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Stdout = a.outStream
	cmd.Stderr = a.errStream

	return cmd.Run()
}
