package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/aaronland/go-aws/v3/ec2"
)

func main() {

	var aws_uri string
	var key string
	var value string

	flag.StringVar(&aws_uri, "aws-uri", "", "...")
	flag.StringVar(&key, "tag-key", "Name", "...")
	flag.StringVar(&value, "tag-value", "", "...")

	flag.Parse()

	ctx := context.Background()

	cl, err := ec2.NewClient(ctx, aws_uri)

	if err != nil {
		log.Fatalf("Failed to create EC2 client, %w", err)
	}

	addrs, err := ec2.GetPublicIPsWithTag(ctx, cl, key, value)

	if err != nil {
		log.Fatalf("Failed to derive public IPs for tag, %w", err)
	}

	for _, a := range addrs {
		fmt.Println(a)
	}

}
