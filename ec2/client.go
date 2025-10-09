package ec2

import (
	"context"

	"github.com/aaronland/go-aws/v3/auth"
	aws_ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func NewClient(ctx context.Context, uri string) (*aws_ec2.Client, error) {

	cfg, err := auth.NewConfig(ctx, uri)

	if err != nil {
		return nil, err
	}

	return aws_ec2.NewFromConfig(cfg), nil

}
