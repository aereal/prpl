package prpl

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"unicode"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/rs/zerolog/log"
)

func NewFromConfig(cfg aws.Config) *Exporter {
	client := ssm.NewFromConfig(cfg)
	return &Exporter{client: client}
}

type Exporter struct {
	client ssm.GetParametersByPathAPIClient
}

func (e *Exporter) ExportParameters(ctx context.Context, paramPath string, environ *[]string) error {
	input := &ssm.GetParametersByPathInput{Path: &paramPath, Recursive: true, WithDecryption: true}
	for {
		var nt string
		if input.NextToken != nil {
			nt = *input.NextToken
		}
		log.Debug().Str("path", *input.Path).Str("nextToken", nt).Msg("GetParametersByPath")
		out, err := e.client.GetParametersByPath(ctx, input)
		if err != nil {
			return fmt.Errorf("failed to GetParametersByPath(nextToken=%s): %w", nt, err)
		}
		log.Debug().Int("parametersCount", len(out.Parameters)).Send()
		for _, param := range out.Parameters {
			if param.Name == nil {
				continue
			}
			if param.Value == nil {
				continue
			}
			name := envName(*param.Name, paramPath)
			*environ = append(*environ, formatEnviron(name, *param.Value))
		}
		if out.NextToken == nil || *out.NextToken == "" {
			break
		}
		input.NextToken = out.NextToken
	}
	return nil
}

func formatEnviron(name, value string) string {
	return fmt.Sprintf("%s=%s", name, value)
}

func envName(path, prefix string) string {
	buf := new(bytes.Buffer)
	var visitedSlash bool
	for _, r := range strings.ToUpper(strings.Replace(path, prefix, "", 1)) {
		if !visitedSlash {
			visitedSlash = true
			continue
		}
		switch {
		case unicode.IsPunct(r):
			buf.WriteByte('_')
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}
