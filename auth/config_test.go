package auth

import (
	"context"
	"testing"
)

func TestNewConfig(t *testing.T) {

	uris := []string{
		"aws://us-east-1?credentials=anon:",
		"aws://?region=us-east-1&credentials=anon:",
	}

	ctx := context.Background()

	for _, uri := range uris {

		_, err := NewConfig(ctx, uri)

		if err != nil {
			t.Fatalf("Unable to create config with URI '%s', %v", uri, err)
		}
	}
}

func TestNewConfigWithCredentials(t *testing.T) {

	creds := []string{
		"anon:",
		"env:",
		"iam:",
		"static:key:secret:",
		"static:key:secret:token",
		"fixtures/credentials:default",
		"fixtures/credentials:example",
		"default",
	}

	ctx := context.Background()

	for _, str := range creds {

		_, err := NewConfigWithCredentialsString(ctx, str)

		if err != nil {
			t.Fatalf("Unable to create config with credentials '%s', %v", str, err)
		}
	}
}
