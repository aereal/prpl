package cli

import (
	"bytes"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestCLI(t *testing.T) {
	testCases := []struct {
		name       string
		args       []string
		wantStatus int
	}{
		{
			"no path",
			[]string{"prpl"},
			1,
		},
		{
			"no command",
			[]string{"prpl", "-path", "/app"},
			1,
		},
		{
			"no command with -debug",
			[]string{"prpl", "-debug", "-path", "/app"},
			1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logStream := new(bytes.Buffer)
			orig := log.Logger
			log.Logger = orig.Output(logStream)
			defer func() {
				log.Logger = orig
			}()

			errStream := new(bytes.Buffer)
			outStream := new(bytes.Buffer)
			app := &App{errStream: errStream, outStream: outStream}
			gotStatus := app.Run(tc.args)
			if gotStatus != tc.wantStatus {
				t.Errorf("status: want=%d got=%d", tc.wantStatus, gotStatus)
			}
			t.Logf("log=%s", logStream.String())
			t.Logf("err=%s", errStream.String())
			t.Logf("out=%s", outStream.String())
		})
	}
}
