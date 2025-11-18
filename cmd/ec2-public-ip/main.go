package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/aaronland/go-aws/v3/ec2"
)

func main() {

	var aws_uri string
	var key string
	var value string

	var verbose bool

	flag.StringVar(&aws_uri, "aws-uri", "", "A valid URI which can be parsed using the `aaronland/go-aws/v3/auth.NewConfig` method.")
	flag.StringVar(&key, "tag-key", "Name", "The name of the tag to filter on.")
	flag.StringVar(&value, "tag-value", "", "The value that the tag (matching -tag-key) should contain")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose (debug) logging.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "List the public IP addresses for EC2 instances whose tag named '-tag-key' contains the string '-tag-value'.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t%s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	cl, err := ec2.NewClient(ctx, aws_uri)

	if err != nil {
		log.Fatalf("Failed to create EC2 client, %v", err)
	}

	addrs, err := ec2.GetPublicIPsWithTag(ctx, cl, key, value)

	if err != nil {
		log.Fatalf("Failed to derive public IPs for tag, %v", err)
	}

	for _, a := range addrs {
		fmt.Println(a)
	}

}
