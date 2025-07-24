package dynamodb

import (
	"context"
	"flag"
	"log/slog"
	"testing"

	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client_uri = flag.String("client-uri", "", "...")

func TestClient(t *testing.T) {

	slog.SetLogLoggerLevel(slog.LevelDebug)

	if *client_uri == "" {
		slog.Info("-client-uri flag not set, skipping test.")
		t.Skip()
	}

	ctx := context.Background()

	client, err := NewClient(ctx, *client_uri)

	if err != nil {
		t.Fatalf("Failed to create new client, %v", err)
	}

	input := &aws_dynamodb.ListTablesInput{}

	rsp, err := client.ListTables(ctx, input)

	if err != nil {
		t.Fatalf("Failed to list tables, %v", err)
	}

	for _, t := range rsp.TableNames {
		slog.Debug(t)
	}

}
