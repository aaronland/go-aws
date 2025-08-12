package main

// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts#Client.AssumeRole
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts#AssumeRoleInput
// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/sts#AssumeRoleOutput

/*

Assume a role with a "trust policy" like this:

{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Statement1",
            "Effect": "Allow",
            "Principal": {
                "AWS": "arn:aws:iam::{AWS_ACCOUNT}:user/{IAM_USER}"
            },
            "Action": "sts:AssumeRole",
            "Condition": {
                "Bool": {
                    "aws:MultiFactorAuthPresent": true
                }
            }
        }
    ]
}
*/

/*

$> ./bin/aws-sts-session -config-uri 'aws://?region={REGION}&credentials={CREDENTIALS}' \
	-role-arn 'arn:aws:iam::{AWS_ACCOUNT}:role/{IAM_ROLE}' \
	-role-session debug \
	-mfa-serial-number arn:aws:iam::{AWS_ACCOUNT}:mfa/{MFA_LABEL} \
	-mfa-token {TOKEN} \
	-session-profile test

2024/11/08 08:23:25 Assumed role "arn:aws:sts::{AWS_ACCOUNT}:assumed-role/{IAM_ROLE}/debug", expires 2024-11-08 17:23:25 +0000 UTC

*/

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aaronland/go-aws/v3/auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func main() {

	var config_uri string

	var role_arn string
	var role_session string
	var role_duration int // update to use ISO8601 duration string

	var mfa_require bool
	var mfa_serial string
	var mfa_token string

	var session_profile string

	flag.StringVar(&config_uri, "config-uri", "", "A valid aaronland/gp-aws-auth.Config URI.")

	flag.StringVar(&role_arn, "role-arn", "", "The AWS role ARN URI of the role you want to assume.")
	flag.StringVar(&role_session, "role-session", "", "A unique name to identify the session.")
	flag.IntVar(&role_duration, "role-duration", 3600, "The duration, in seconds, of the role session.")

	flag.BoolVar(&mfa_require, "mfa", true, "Require a valid MFA token code when assuming role.")
	flag.StringVar(&mfa_serial, "mfa-serial-number", "", "The unique identifier of the MFA device being used for authentication.")
	flag.StringVar(&mfa_token, "mfa-token", "", "A valid MFA token string. If empty then data will be read from a command line prompt.")

	flag.StringVar(&session_profile, "session-profile", "", "The name of the AWS credentials profile to associate the temporary credentials with.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Generate STS credentials for a given profile and MFA token and then write those credentials back to an AWS \"credentials\" file in a specific profile section.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Valid options are:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	ctx := context.Background()

	cfg, err := auth.NewConfig(ctx, config_uri)

	if err != nil {
		log.Fatalf("Failed to create new config, %v", err)
	}

	creds, err := auth.NewCredentials()

	if err != nil {
		log.Fatalf("Failed to create new credentials, %v", err)
	}

	cl := sts.NewFromConfig(cfg)

	assume_opts := &sts.AssumeRoleInput{
		RoleArn:         aws.String(role_arn),
		RoleSessionName: aws.String(role_session),
		DurationSeconds: aws.Int32(int32(role_duration)),
	}

	if mfa_require {

		mfa_token = strings.TrimSpace(mfa_token)

		if mfa_token == "" {
			mfa_token = readline("Enter your MFA token code:")
		}

		assume_opts.SerialNumber = aws.String(mfa_serial)
		assume_opts.TokenCode = aws.String(mfa_token)
	}

	rsp, err := cl.AssumeRole(ctx, assume_opts)

	if err != nil {
		log.Fatalf("Failed to assume role, %v", err)
	}

	session_creds := rsp.Credentials

	err = creds.SetSessionCredentialsWithProfile(ctx, session_profile, session_creds)

	if err != nil {
		log.Fatalf("Failed to get credentials with session profile, %v", err)
	}

	log.Printf(`Assumed role "%s", expires %v`, *rsp.AssumedRoleUser.Arn, *session_creds.Expiration)
}

func readline(prompt string) string {

	var input string

	fmt.Print(fmt.Sprintf("%s ", prompt))
	fmt.Scanf("%s", &input)

	// go-sanitize strings here? (20180621/thisisaaronland)

	return strings.Trim(input, " ")
}
