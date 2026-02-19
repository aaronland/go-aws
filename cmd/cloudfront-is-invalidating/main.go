package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/cloudfront"
)

func main() {

	var client_uri string
	var distribution_id string

	flag.StringVar(&client_uri, "client-uri", "", "A valid client URI in the form of 'aws://?region={AWS_REGION}&credentials={CREDENTIALS}' where '{CREDENTIAL}' is expected to be a valid aaronland/go-aws-auth credential string.")
	flag.StringVar(&distribution_id, "distribution-id", "", "A valid AWS CloudFront distribution ID.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Check whether a CloudFront distribution is running an invalidation task.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s uri(N) uri(N)\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	cl, err := cloudfront.NewClientWithURI(ctx, client_uri)

	if err != nil {
		log.Fatalf("Failed to create client, %v", err)
	}

	is_running, err := cloudfront.IsInvalidationRunning(ctx, cl, distribution_id)

	if err != nil {
		log.Fatalf("Failed to invalidate paths, %v", err)
	}

	fmt.Printf("%v\n", is_running)
}
