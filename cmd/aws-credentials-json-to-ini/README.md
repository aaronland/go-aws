# aws-credentials-json-to-ini

`aws-credentials-json-to-ini` reads JSON-encoded AWS credentials information and generates an AWS ini-style configuration file with those data.

```
$> ./bin/aws-credentials-json-to-ini -h
Usage of ./bin/aws-credentials-json-to-ini:
  -ini string
    	Path to the ini-style file where AWS credentials should be written. If "-" then data will be written to STDOUT.
  -json string
    	Path to the JSON file containing AWS credentials. If "-" then data will be read from STDIN.
  -name string
    	The name of the ini section where AWS credentials should be written. (default "default")
  -region string
    	The AWS region for the AWS credentials. (default "us-east-1")
```

For example:

```
$> go bin/aws-cognito-credentials \
	-aws-config-uri 'aws://us-east-1?credentials=session' \
	-identity-pool-id us-east-1:{GUID} \
	-login org.sfomuseum=bob
	-role-session-name bob -role-arn 'arn:aws:iam::{ACCOUNT_ID}:role/{ROLE}' \

| ./bin/aws-credentials-json-to-ini -json - -ini -

[default]
region = us-east-1
aws_access_key_id = ...
aws_secret_access_key = ...
aws_session_token = ...
```