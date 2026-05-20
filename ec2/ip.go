package ec2

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"

	aws_ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
)

// GetPublicIPsWithTags returns the public IP addresses for EC2 instances whose tag named 'key' contains the string 'value'.
// Instances whose key=value tag match but don't have a public IP address are logged as a warning.
func GetPublicIPsWithTag(ctx context.Context, cl *aws_ec2.Client, key string, value string) ([]string, error) {

	re_str := fmt.Sprintf(`(^|\s)%s(\s|$)`, value)
	re_value, err := regexp.Compile(re_str)

	if err != nil {
		return nil, fmt.Errorf("Failed to compile regexp for value, %w", err)
	}

	addrs := make([]string, 0)

	input := &aws_ec2.DescribeInstancesInput{}

	out, err := cl.DescribeInstances(ctx, input)

	if err != nil {
		return nil, fmt.Errorf("DescribeInstances failed: %w", err)
	}

	for _, rsv := range out.Reservations {

		for _, inst := range rsv.Instances {

			logger := slog.Default()
			logger = logger.With("instance ID", *inst.InstanceId)

			for _, t := range inst.Tags {

				logger.Debug("Tag", "key", *t.Key, "value", *t.Value)

				if *t.Key != key {
					continue
				}

				if !re_value.MatchString(*t.Value) {
					continue
				}

				if inst.PublicIpAddress == nil || *inst.PublicIpAddress == "" {
					logger.Warn("Instance missing public IP address", "tag", *t.Value)
					continue
				}

				addrs = append(addrs, *inst.PublicIpAddress)
			}
		}
	}

	return addrs, nil
}
