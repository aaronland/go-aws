//go:build local

package dynamodb

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	aws_dynamodb_types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestLocalClient(t *testing.T) {

	slog.SetLogLoggerLevel(slog.LevelDebug)
	ctx := context.Background()

	client, err := NewClient(ctx, LOCAL_CLIENT_URI)

	if err != nil {
		t.Fatalf("Failed to create new client, %v", err)
	}

	test_table := "test"

	table_opts := &CreateTablesOptions{
		Tables: map[string]*aws_dynamodb.CreateTableInput{
			test_table: &aws_dynamodb.CreateTableInput{
				KeySchema: []aws_dynamodb_types.KeySchemaElement{
					{
						AttributeName: aws.String("Code"),
						KeyType:       "HASH", // partition key
					},
				},
				AttributeDefinitions: []aws_dynamodb_types.AttributeDefinition{
					{
						AttributeName: aws.String("Code"),
						AttributeType: "S",
					},
				},
				BillingMode: aws_dynamodb_types.BillingModePayPerRequest,
			},
		},
	}

	err = CreateTables(ctx, client, table_opts)

	if err != nil {
		t.Fatalf("Failed to create tables, %v", err)
	}

	list_opts := &aws_dynamodb.ListTablesInput{}

	list_rsp, err := client.ListTables(ctx, list_opts)

	if err != nil {
		t.Fatalf("Failed to list tables, %v", err)
	}

	for _, t := range list_rsp.TableNames {
		slog.Debug(t)
	}

	delete_opts := &aws_dynamodb.DeleteTableInput{
		TableName: aws.String(test_table),
	}

	_, err = client.DeleteTable(ctx, delete_opts)

	if err != nil {
		t.Fatalf("Failed to delete table, %v", err)
	}

}
