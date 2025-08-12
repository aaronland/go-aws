# aws-set-env

`aws-set-env` is a command line tool to assign required AWS authentication environment variables for a given profile in a AWS .credentials file.

```
$> ./bin/aws-set-env -h
Usage of ./bin/aws-set-env:
  -profile string
    	A valid AWS credentials profile (default "default")
  -session-token
    	Require AWS_SESSION_TOKEN environment variable (default true)
```
