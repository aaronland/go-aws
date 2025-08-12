# aws-sign-request

`aws-sign-request` signs a HTTP request with an AWS "v4" signature, optionally executing the request and emitting the output to STDOUT or writing the request itself to STDOUT.

```
$> ./bin/aws-sign-request -h
Usage of ./bin/aws-sign-request:
  -api-signing-name string
    	The name the API uses to identify the service the request is scoped to.
  -api-signing-region string
    	If empty then the value of the region associated with the AWS config/credentials will be used.
  -credentials-uri string
    	A valid aaronland/go-aws-auth config URI.
  -debug
    	Enable verbose debug logging to STDOUT.	
  -do
    	If true then execute the signed request and output the response to STDOUT.
  -header value
    	Zero or more HTTP headers to assign to the request in the form of key=value.
  -method string
    	A valid HTTP method. (default "GET")
  -uri string
    	The URI you are trying to sign.
```

For example, to call a Lambda Function URL:

```
$> bin/aws-sign-request \
	-credentials-uri 'aws://{REGION}?credentials=iam:' \
	-api-signing-name 'lambda' \
	-uri https://{GIBBERISH}.lambda-url.{REGION}.on.aws/api/point-in-polygon \
	-method POST \
	-do \
	'{"latitude": 25.0, "longitude": -45.6 }' \
	
	| jq

{
  "places": [
    {
      "wof:id": "404528709",
      "wof:parent_id": "-1",
      "wof:name": "North Atlantic Ocean",
      "wof:country": "",
      "wof:placetype": "ocean",
      "mz:latitude": 0,
      "mz:longitude": 0,
      "mz:min_latitude": 24.965357,
      "mz:min_longitude": 0,
      "mz:max_latitude": -45.616087,
      "mz:max_longitude": -45.570425,
      "mz:is_current": 1,
      "mz:is_deprecated": -1,
      "mz:is_ceased": -1,
      "mz:is_superseded": 0,
      "mz:is_superseding": 0,
      "edtf:inception": "",
      "edtf:cessation": "",
      "wof:supersedes": [],
      "wof:superseded_by": [],
      "wof:belongsto": [],
      "wof:path": "404/528/709/404528709.geojson",
      "wof:repo": "whosonfirst-data-admin-xy",
      "wof:lastmodified": 1690923898
    }
  ]
}
```
