package acm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_acm "github.com/aws/aws-sdk-go-v2/service/acm"
)

func ExportCertificate(ctx context.Context, cl *aws_acm.Client, arn string, pswd string) (*Certificate, error) {

	opts := &aws_acm.ExportCertificateInput{
		CertificateArn: aws.String(arn),
		Passphrase:     []byte(pswd),
	}

	rsp, err := cl.ExportCertificate(ctx, opts)

	if err != nil {
		return nil, err
	}

	cert := &Certificate{
		Certificate:      []byte(*rsp.Certificate),
		CertificateChain: []byte(*rsp.CertificateChain),
		PrivateKey:       []byte(*rsp.PrivateKey),
	}

	return cert, nil
}
