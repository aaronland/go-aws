// aws-sign-request signs a HTTP request with an AWS "v4" signature, optionally executing the
// request and emitting the output to STDOUT or writing the request itself to STDOUT.
package main

import (
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aaronland/go-aws/v3/auth"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/logging"
	"github.com/sfomuseum/go-flags/multi"
)

func main() {

	var headers multi.KeyValueString

	var api_signing_name string
	var api_signing_region string

	var credentials_uri string
	var method string
	var uri string
	var do bool
	var debug bool

	flag.StringVar(&api_signing_name, "api-signing-name", "", "The name the API uses to identify the service the request is scoped to.")
	flag.StringVar(&api_signing_region, "api-signing-region", "", "If empty then the value of the region associated with the AWS config/credentials will be used.")

	flag.StringVar(&method, "method", "GET", "A valid HTTP method.")
	flag.StringVar(&uri, "uri", "", "The URI you are trying to sign.")
	flag.StringVar(&credentials_uri, "credentials-uri", "", "A valid aaronland/go-aws-auth config URI.")
	flag.BoolVar(&do, "do", false, "If true then execute the signed request and output the response to STDOUT.")
	flag.Var(&headers, "header", "Zero or more HTTP headers to assign to the request in the form of key=value.")
	flag.BoolVar(&debug, "debug", false, "Enable verbose debug logging to STDOUT.")

	flag.Parse()

	ctx := context.Background()

	body_r := strings.NewReader(strings.Join(flag.Args(), " "))

	//

	cfg, err := auth.NewConfig(ctx, credentials_uri)

	if err != nil {
		log.Fatalf("Failed to create new config, %v", err)
	}

	creds, err := cfg.Credentials.Retrieve(ctx)

	if err != nil {
		log.Fatalf("Failed to derive credentials from config, %v", err)
	}

	if api_signing_region == "" {
		api_signing_region = cfg.Region
	}

	// https://github.com/aws/aws-sdk-go-v2/blob/main/aws/signer/v4/v4.go#L287

	h := sha256.New()

	_, err = io.Copy(h, body_r)

	if err != nil {
		log.Fatalf("Failed to hash request body, %v", err)
	}

	body_sha256 := fmt.Sprintf("%x", h.Sum(nil))

	_, err = body_r.Seek(0, 0)

	if err != nil {
		log.Fatalf("Failed to rewind message body, %v", err)
	}

	req, err := http.NewRequest(method, uri, body_r)

	if err != nil {
		log.Fatalf("Failed to create new HTTP request, %v", err)
	}

	for _, h := range headers {
		req.Header.Set(h.Key(), h.Value().(string))
	}

	signer := v4.NewSigner(func(opts *v4.SignerOptions) {
		if debug {
			opts.LogSigning = true
			opts.Logger = logging.NewStandardLogger(os.Stdout)
		}
	})

	err = signer.SignHTTP(ctx, creds, req, body_sha256, api_signing_name, api_signing_region, time.Now())

	if err != nil {
		log.Fatalf("Failed to sign request, %v", err)
	}

	if !do {

		err = req.Write(os.Stdout)

		if err != nil {
			log.Fatalf("Failed to write request, %v", err)
		}

		return
	}

	cl := http.Client{}
	rsp, err := cl.Do(req)

	if err != nil {
		log.Fatalf("Failed to execute request, %v", err)
	}

	defer rsp.Body.Close()

	_, err = io.Copy(os.Stdout, rsp.Body)

	if err != nil {
		log.Fatalf("Failed to read response, %v", err)
	}

}
