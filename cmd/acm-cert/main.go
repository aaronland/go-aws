package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/acm"
)

func main() {

	var client_uri string
	var arn string

	flag.StringVar(&client_uri, "client-uri", "", "...")
	flag.StringVar(&arn, "arn", "", "...")

	flag.Parse()

	ctx := context.Background()

	cl, err := acm.NewClient(ctx, client_uri)

	if err != nil {
		log.Fatal(err)
	}

	rsp, err := acm.ExportCertificate(ctx, cl, arn)

	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(rsp)

	if err != nil {
		log.Fatal(err)
	}
}
