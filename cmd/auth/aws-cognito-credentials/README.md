# aws-cognito-credentials

`aws-cognito-credentials` generates temporary STS credentials for a given user in a Cognito identity pool.

```
$> ./bin/aws-cognito-credentials -h
Usage of ./bin/aws-cognito-credentials:
  -aws-config-uri string
    	A valid github.com/aaronland/go-aws-auth.Config URI.
  -duration int
    	The duration, in seconds, of the role session. Can not be less than 900. (default 900)
  -identity-pool-id string
    	A valid AWS Cognito Identity Pool ID.
  -login value
    	One or more key=value strings mapping to AWS Cognito authentication providers.
  -role-arn string
    	A valid AWS IAM role ARN to assign to STS credentials.
  -role-session-name string
    	An identifier for the assumed role session.
  -session-policy value
    	Zero or more IAM ARNs to use as session policies to supplement the default role ARN.	
```

For example:

```
$> go bin/aws-cognito-credentials \
	-aws-config-uri 'aws://us-east-1?credentials=session' \
	-identity-pool-id us-east-1:{GUID} \
	-login org.sfomuseum=bob
	-role-session-name bob -role-arn 'arn:aws:iam::{ACCOUNT_ID}:role/{ROLE}' \
	
| jq
	
{
  "AccessKeyId": "...",
  "Expiration": "...",
  "SecretAccessKey": "...",
  "SessionToken": "..."
}
```
