package ec2

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func GetPublicIPsWithTag(ctx context.Context, cl *aws_ec2.Client, key string, value string) ([]string, error) {

	addrs := make([]string, 0)

	slog.Info("Filter", "key", key, "value", value)

	filters := []types.Filter{
		{
			Name:   aws.String(fmt.Sprintf("tag:%s", key)),
			Values: []string{value},
		},
	}

	slog.Info("DEBUG", "filters", filters)

	input := &aws_ec2.DescribeInstancesInput{
		//Filters: filters,
	}

	out, err := cl.DescribeInstances(ctx, input)

	if err != nil {
		return nil, fmt.Errorf("DescribeInstances failed: %w", err)
	}

	// Walk the reservations → instances looking for a public IP.
	for _, rsv := range out.Reservations {

		for _, inst := range rsv.Instances {

			logger := slog.Default()
			logger = logger.With("instance ID", *inst.InstanceId)

			if inst.PublicIpAddress == nil || *inst.PublicIpAddress == "" {
				logger.Debug("No address, skip")
				continue
			}

			for _, t := range inst.Tags {

				logger.Debug("Tag", "key", *t.Key, "value", *t.Value)

				if *t.Key != key {
					continue
				}

				if strings.Contains(*t.Value, value) {
					addrs = append(addrs, *inst.PublicIpAddress)
				}
			}
		}
	}

	return addrs, nil
}
