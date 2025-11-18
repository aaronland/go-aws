package cloudfront

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

// GetDistributionIDByHostname queries CloudFront for all distributions and returns
// the Distribution ID whose Alias list contains the supplied hostname.
func GetDistributionIDByHostname(ctx context.Context, client *cloudfront.Client, hostname string) (string, error) {

	hostname = strings.ToLower(strings.TrimSpace(hostname))

	paginator := cloudfront.NewListDistributionsPaginator(client, &cloudfront.ListDistributionsInput{})

	for paginator.HasMorePages() {

		page, err := paginator.NextPage(ctx)

		if err != nil {
			return "", fmt.Errorf("cloudfront ListDistributions: %w", err)
		}

		for _, d := range page.DistributionList.Items {
			if d.Aliases == nil || len(d.Aliases.Items) == 0 {
				// This distribution has no custom domain names – skip it.
				continue
			}

			for _, a := range d.Aliases.Items {
				if strings.EqualFold(a, hostname) {
					return aws.ToString(d.Id), nil
				}
			}
		}
	}

	return "", fmt.Errorf("no CloudFront distribution found for hostname %q", hostname)
}
