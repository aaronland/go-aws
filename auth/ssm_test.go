package auth

import (
	"context"
	_ "fmt"
	"testing"
)

func TestNewSSMClient(t *testing.T) {

	uris := []string{
		"aws://us-east-1?credentials=anon:",
	}

	ctx := context.Background()

	for _, uri := range uris {

		_, err := NewSSMClient(ctx, uri)

		if err != nil {
			t.Fatalf("Unable to create SSM client with URI '%s', %v", uri, err)
		}
	}
}

func TestNewSSMClientWithCredentials(t *testing.T) {

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

		_, err := NewSSMClientWithCredentialsString(ctx, str)

		if err != nil {
			t.Fatalf("Unable to create config with credentials '%s', %v", str, err)
		}
	}
}
