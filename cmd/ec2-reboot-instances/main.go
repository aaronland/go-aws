package main

import (
	"context"
	"log"

	"github.com/aaronland/go-aws/v3/ec2"
	"github.com/sfomuseum/go-flags/flagset"
)

func main() {

	var client_uri string

	fs := flagset.NewFlagSet("reboot")
	fs.StringVar(&client_uri, "client-uri", "", "A valid aaronland/go-aws/auth config URI in the form of aws://?region={REGION}&credentials={CREDENTIALS}")

	flagset.Parse(fs)

	ctx := context.Background()

	instances := fs.Args()

	err := ec2.RebootInstances(ctx, client_uri, instances...)

	if err != nil {
		log.Fatalf("Failed to reboot instances, %v", err)
	}
}
