# invalidate

Invalidate one or more URIs from a CloudFront distribution.

```
$> ./bin/invalidate -h
Invalidate one or more URIs from a CloudFront distribution.
Usage:
	 ./bin/invalidate uri(N) uri(N)
  -client-uri string
    	A valid client URI in the form of 'aws://?region={AWS_REGION}&credentials={CREDENTIALS}' where '{CREDENTIAL}' is expected to be a valid aaronland/go-aws-auth credential string.
  -distribution-id string
    	A valid AWS CloudFront distribution ID.
```
