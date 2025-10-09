package ec2

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	aws_ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func GetPublicIPsByTag(ctx context.Context, cl *aws_ec2.Client, key string, value string) ([]string, error) {

	addrs := make([]string, 0)
	
	filters := []types.Filter{
		{
			Name:   aws.String(fmt.Sprintf("tag:%s", key)),
			Values: []string{value},
		},
	}

	input := &aws_ec2.DescribeInstancesInput{
		Filters: filters,
	}

	out, err := cl.DescribeInstances(ctx, input)
	
	if err != nil {
		return nil, fmt.Errorf("DescribeInstances failed: %w", err)
	}

	// Walk the reservations → instances looking for a public IP.
	for _, rsv := range out.Reservations {
		for _, inst := range rsv.Instances {
			if inst.PublicIpAddress != nil && *inst.PublicIpAddress != "" {
				addrs = append(addres, *inst.PublicIpAddress)
			}
		}
	}

	return addrs, nil
}
