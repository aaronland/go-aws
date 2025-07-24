package cloudfront

import (
	"context"
	"fmt"

	"github.com/aaronland/go-aws/v3/auth"
	aws_cloudfront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

// NewClientWithURI will return a new `aws-sdk-go/service/cloudfront.CloudFront` instance
// derived from 'uri' which is expected to take the form of:
//
//	aws://?region={AWS_REGION}&credentials={CREDENTIALS}
//
// Where '{CREDENTIALS}' is expected to be a valid `aaronland/go-aws-session` credentials
// string.
func NewClientWithURI(ctx context.Context, uri string) (*aws_cloudfront.Client, error) {

	cfg, err := auth.NewConfig(ctx, uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to derive config, %w", err)
	}

	return aws_cloudfront.NewFromConfig(cfg), nil
}
