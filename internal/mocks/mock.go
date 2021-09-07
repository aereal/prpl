//go:generate go run github.com/golang/mock/mockgen -package mocks -destination ./mock_ssm.go github.com/aws/aws-sdk-go-v2/service/ssm GetParametersByPathAPIClient

package mocks
