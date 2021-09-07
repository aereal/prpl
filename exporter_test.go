package prpl

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aereal/prpl/internal/mocks"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestExportParameters(t *testing.T) {
	testCases := []struct {
		name         string
		paramPath    string
		mockedOutput *ssm.GetParametersByPathOutput
		mockedErr    error
		wantEnv      []string
	}{
		{
			"empty",
			"/my-app/staging",
			&ssm.GetParametersByPathOutput{Parameters: []types.Parameter{}},
			nil,
			[]string(nil),
		},
		{
			"some",
			"/my-app/staging",
			&ssm.GetParametersByPathOutput{
				Parameters: []types.Parameter{
					stringParam("/my-app/staging/creds/id", "my-id"),
				},
			},
			nil,
			[]string{"CREDS_ID=my-id"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			m := mocks.NewMockGetParametersByPathAPIClient(ctrl)
			m.EXPECT().GetParametersByPath(gomock.Any(), gomock.Any()).Times(1).Return(tc.mockedOutput, tc.mockedErr)
			exporter := &Exporter{client: m}
			ctx := context.Background()
			var gotEnv []string
			if err := exporter.ExportParameters(ctx, tc.paramPath, &gotEnv); err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(gotEnv, tc.wantEnv); diff != "" {
				t.Errorf("env (-got, +want):\n%s", diff)
			}
		})
	}
}

func paramARN(path string) arn.ARN {
	return arn.ARN{
		Region:    "us-east-1",
		AccountID: "1234567890",
		Service:   "ssm",
		Resource:  fmt.Sprintf("resource/%s", path),
	}
}

func stringParam(path string, value string) types.Parameter {
	return types.Parameter{
		ARN:              aws.String(paramARN(path).String()),
		Name:             aws.String(path),
		Type:             types.ParameterTypeString,
		DataType:         aws.String("text"),
		LastModifiedDate: aws.Time(time.Now()),
		Value:            &value,
	}
}
