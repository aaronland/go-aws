# aws-imds-credentials

`aws-imds-credentials` returns the current `aws.Credentials` derived from the EC2 IMDS API. For example:

```
$> ./bin/aws-imds-credentials | jq
{
  "AccessKeyID": "...",
  "SecretAccessKey": "...",
  "SessionToken": "...",
  "Source": "EC2RoleProvider",
  "CanExpire": true,
  "Expires": "2024-03-28T19:44:42.59621653Z"
}
```
