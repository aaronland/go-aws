package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/acm"
	"github.com/aaronland/go-string/random"
)

func main() {

	var client_uri string
	var password string
	var arn string
	var remove_password bool

	flag.StringVar(&client_uri, "client-uri", "", "...")
	flag.StringVar(&password, "password", "", "...")
	flag.StringVar(&arn, "arn", "", "...")
	flag.BoolVar(&remove_password, "remove-password", false, "...")

	flag.Parse()

	ctx := context.Background()

	if password == "" {

		v, err := randomPassword()

		if err != nil {
			log.Fatal(err)
		}

		password = v
	}

	cl, err := acm.NewClient(ctx, client_uri)

	if err != nil {
		log.Fatal(err)
	}

	cert, err := acm.ExportCertificate(ctx, cl, arn, password)

	if err != nil {
		log.Fatal(err)
	}

	if remove_password {

		key_nopass, err := cert.RemovePassword(password)

		if err != nil {
			log.Fatal(err)
		}

		cert.PrivateKey = key_nopass
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(cert)

	if err != nil {
		log.Fatal(err)
	}

}

func randomPassword() (string, error) {

	opts := &random.Options{
		Length:       32,
		AlphaNumeric: true,
	}

	return random.String(opts)
}
