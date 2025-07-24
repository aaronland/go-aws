package cloudfront

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_cloudfront "github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

// InvalidatePaths will issue a "CreateInvalidation" request for 'uris' in 'distribution_id'. It will return the
// invalidation ID and caller reference associated with the request.
func InvalidatePaths(ctx context.Context, cl *aws_cloudfront.Client, distribution_id string, uris ...string) (string, string, error) {

	count := len(uris)

	paths := &types.Paths{
		Items:    uris,
		Quantity: aws.Int32(int32(count)),
	}

	enc_items, err := json.Marshal(uris)

	if err != nil {
		return "", "", fmt.Errorf("Failed to encode items used to derive reference, %v", err)
	}

	sum := sha256.Sum256(enc_items)
	ref := fmt.Sprintf("%x", sum)

	batch := &types.InvalidationBatch{
		CallerReference: aws.String(ref),
		Paths:           paths,
	}

	input := &aws_cloudfront.CreateInvalidationInput{
		DistributionId:    aws.String(distribution_id),
		InvalidationBatch: batch,
	}

	rsp, err := cl.CreateInvalidation(ctx, input)

	if err != nil {
		return "", "", fmt.Errorf("Failed to create invalidation, %v", err)
	}

	id := *rsp.Invalidation.Id

	return id, ref, nil
}
