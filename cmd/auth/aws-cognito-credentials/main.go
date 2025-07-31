// aws-cognito-credentials generates temporary STS credentials for a given user in a Cognito identity pool.
package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/auth"
	"github.com/sfomuseum/go-flags/multi"
)

func main() {

	var aws_config_uri string

	var identity_pool_id string
	var role_arn string
	var role_session_name string
	var duration int

	var kv_logins multi.KeyValueString
	var session_policies multi.MultiString

	flag.StringVar(&aws_config_uri, "aws-config-uri", "", "A valid github.com/aaronland/go-aws-auth.Config URI.")

	flag.StringVar(&identity_pool_id, "identity-pool-id", "", "A valid AWS Cognito Identity Pool ID.")
	flag.StringVar(&role_arn, "role-arn", "", "A valid AWS IAM role ARN to assign to STS credentials.")
	flag.StringVar(&role_session_name, "role-session-name", "", "An identifier for the assumed role session.")
	flag.IntVar(&duration, "duration", 900, "The duration, in seconds, of the role session. Can not be less than 900.") // Note: Can not be less than 900
	flag.Var(&kv_logins, "login", "One or more key=value strings mapping to AWS Cognito authentication providers.")
	flag.Var(&session_policies, "session-policy", "Zero or more IAM ARNs to use as session policies to supplement the default role ARN.")

	flag.Parse()

	ctx := context.Background()

	cfg, err := auth.NewConfig(ctx, aws_config_uri)

	if err != nil {
		log.Fatalf("Failed to derive AWS config, %v", err)
	}

	logins := make(map[string]string, 0)

	for _, kv := range kv_logins {
		logins[kv.Key()] = kv.Value().(string)
	}

	opts := &auth.STSCredentialsForDeveloperIdentityOptions{
		RoleArn:         role_arn,
		RoleSessionName: role_session_name,
		Duration:        int32(duration),
		IdentityPoolId:  identity_pool_id,
		Logins:          logins,
		Policies:        session_policies,
	}

	creds, err := auth.STSCredentialsForDeveloperIdentity(ctx, cfg, opts)

	if err != nil {
		log.Fatalf("Failed to derive credentials, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(creds)

	if err != nil {
		log.Fatalf("Failed to encode credentials, %v", err)
	}
}
