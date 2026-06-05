package acm

import (
	"context"

	"github.com/aaronland/go-string/random"
	"github.com/aws/aws-sdk-go-v2/aws"
	aws_acm "github.com/aws/aws-sdk-go-v2/service/acm"
)

func ExportCertificate(ctx context.Context, cl *aws_acm.Client, arn string) (any, error) {

	pswd, err := password()

	if err != nil {
		return nil, err
	}

	opts := &aws_acm.ExportCertificateInput{
		CertificateArn: aws.String(arn),
		Passphrase:     []byte(pswd),
	}

	rsp, err := cl.ExportCertificate(ctx, opts)

	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func password() (string, error) {

	opts := &random.Options{
		Length:       32,
		AlphaNumeric: true,
	}

	return random.String(opts)
}
