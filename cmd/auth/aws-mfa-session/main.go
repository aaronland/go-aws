// aws-mfa-session is a command line to create session-based authentication keys and secrets for
// a given profile and multi-factor authentication (MFA) token and then writing that key and secret
// back to a "credentials" file in a specific profile section. For example, when used in a Makefile with
// https://github.com/Yubico/yubikey-manager/tree/master/ykman
//
//	$(eval CODE := $(shell ykman oath code sfomuseum:aws | awk '{ print $$2 }'))
//	 bin/$aws-mfa-session -code $(CODE) -duration PT8H
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aaronland/go-aws/v3/auth"
	"github.com/sfomuseum/iso8601duration"
)

func readline(prompt string) string {

	var input string

	fmt.Print(fmt.Sprintf("%s ", prompt))
	fmt.Scanf("%s", &input)

	// go-sanitize strings here? (20180621/thisisaaronland)

	return strings.Trim(input, " ")
}

func main() {

	var device string
	var profile string
	var code string
	var session_profile string
	var session_duration string

	flag.StringVar(&profile, "profile", "default", "A valid AWS credentials profile")
	flag.StringVar(&device, "device", "", "The device ID (serial number) used to validate MFA codes")
	flag.StringVar(&code, "code", "", "A valid MFA code. If empty the application will block and prompt the user")
	flag.StringVar(&session_profile, "session-profile", "session", "The name of the AWS credentials profile to update with session credentials")
	flag.StringVar(&session_duration, "duration", "PT1H", "A valid ISO8601 duration string indicating how long the session should last (months are currently not supported)")

	flag.Parse()

	ctx := context.Background()

	d, err := duration.FromString(session_duration)

	if err != nil {
		log.Fatalf("Failed to parse session duration, %v", err)
	}

	ttl := d.ToDuration().Seconds()
	ttl_32 := int32(ttl)

	creds, err := auth.NewCredentials()

	if err != nil {
		log.Fatalf("Failed to create new credentials, %v", err)
	}

	aws_cfg, err := creds.AWSConfigWithProfile(ctx, profile)

	if err != nil {
		log.Fatalf("Failed to create AWS config, %v", err)
	}

	code = strings.TrimSpace(code)

	if code == "" {

		code = readline("Enter your MFA token code:")

		if code == "" {
			log.Fatalf("Missing MFA code")
		}
	}

	session_creds, err := auth.GetCredentialsWithMFAWithContext(ctx, aws_cfg, device, code, ttl_32)

	if err != nil {
		log.Fatalf("Failed to get credentials with MFA, %v", err)
	}

	err = creds.SetSessionCredentialsWithProfile(ctx, session_profile, session_creds)

	if err != nil {
		log.Fatalf("Failed to get credentials with session profile, %v", err)
	}

	now := time.Now()
	then := now.Add(d.ToDuration())

	log.Printf("Updated session credentials for '%s' profile, expires %s (%s)\n", session_profile, then.Format(time.Stamp), *session_creds.Expiration)
}
