package ec2

import (
	"context"
	"fmt"

	"github.com/aaronland/go-aws/v3/auth"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// RebootInstances will reboot the AWS EC2 instances listed in 'instances'.
func RebootInstances(ctx context.Context, client_uri string, instances ...string) error {

	cfg, err := auth.NewConfig(ctx, client_uri)

	if err != nil {
		return err
	}

	ec2Client := ec2.NewFromConfig(cfg)

	input := &ec2.RebootInstancesInput{
		InstanceIds: instances,
	}

	_, err = ec2Client.RebootInstances(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to reboot instances, %w", err)
	}

	return nil
}
