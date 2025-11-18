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
	var hostname string
	
	flag.StringVar(&client_uri, "client-uri", "", "A valid client URI in the form of 'aws://?region={AWS_REGION}&credentials={CREDENTIALS}' where '{CREDENTIAL}' is expected to be a valid aaronland/go-aws-auth credential string.")
	flag.StringVar(&hostname, "hostname", "", "The hostname to derive a CloudFront distibution ID for.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Derive the CloudFront distribution ID for a given hostname.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	cl, err := cloudfront.NewClientWithURI(ctx, client_uri)

	if err != nil {
		log.Fatalf("Failed to create client, %v", err)
	}

	id, err := cloudfront.GetDistributionIDByHostname(ctx, cl, hostname)

	if err != nil {
		log.Fatalf("Failed to derive distribution ID, %v", err)
	}

	fmt.Println(id)
}
