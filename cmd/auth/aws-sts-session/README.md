# aws-sts-session

Generate STS credentials for a given profile and MFA token and then write those credentials back to an AWS "credentials" file in a specific profile section.

```
$> ./bin/aws-sts-session -h
Generate STS credentials for a given profile and MFA token and then write those credentials back to an AWS "credentials" file in a specific profile section.
Usage:
	 ./bin/aws-sts-session [options]
Valid options are:
  -config-uri string
    	A valid aaronland/gp-aws-auth.Config URI.
  -mfa
    	Require a valid MFA token code when assuming role. (default true)
  -mfa-serial-number string
    	The unique identifier of the MFA device being used for authentication.
  -mfa-token string
    	A valid MFA token string. If empty then data will be read from a command line prompt.
  -role-arn string
    	The AWS role ARN URI of the role you want to assume.
  -role-duration int
    	The duration, in seconds, of the role session. (default 3600)
  -role-session string
    	A unique name to identify the session.
  -session-profile string
    	The name of the AWS credentials profile to associate the temporary credentials with.
```

For example:

```
$> ./bin/aws-sts-session -config-uri 'aws://?region={REGION}&credentials={CREDENTIALS}' \
	-role-arn 'arn:aws:iam::{AWS_ACCOUNT}:role/{IAM_ROLE}' \
	-role-session debug \
	-mfa-serial-number arn:aws:iam::{AWS_ACCOUNT}:mfa/{MFA_LABEL} \
	-mfa-token {TOKEN} \
	-session-profile test

2024/11/08 08:23:25 Assumed role "arn:aws:sts::{AWS_ACCOUNT}:assumed-role/{IAM_ROLE}/debug", expires 2024-11-08 17:23:25 +0000 UTC
```

Note that this assumes a role with a "trust policy" equivalent to this:

```
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
```
