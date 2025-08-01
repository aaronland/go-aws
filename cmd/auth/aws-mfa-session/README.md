# aws-mfa-session

`aws-mfa-session` is a command line to create session-based authentication keys and secrets for a given profile and multi-factor authentication (MFA) token and then writing that key and secret back to a "credentials" file in a specific profile section.

```
$> ./bin/aws-mfa-session -h
Usage of ./bin/aws-mfa-session:
  -code string
    	A valid MFA code. If empty the application will block and prompt the user
  -device string
    	The device ID (serial number) used to validate MFA codes
  -duration string
    	A valid ISO8601 duration string indicating how long the session should last (months are currently not supported) (default "PT1H")
  -profile string
    	A valid AWS credentials profile (default "default")
  -session-profile string
    	The name of the AWS credentials profile to update with session credentials (default "session")
```

For example:

```
$> ./bin/aws-mfa-session -profile {PROFILE} -device {MFA_DEVICE_ARN} -duration PT8H
Enter your MFA token code: 123456
2018/07/26 09:47:09 Updated session credentials for 'session' profile, expires Jul 26 17:47:09 (2018-07-27 00:51:52 +0000 UTC)
```

Note that the `-device` flag is optional. If empty and there is only one MFA device registered that one will be used. If empty and there are multiple MFA devices registered then an error will be thrown.
