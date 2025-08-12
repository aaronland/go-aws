package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	aa_dynamodb "github.com/aaronland/go-aws/v3/dynamodb"
	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {

	client_uri := flag.String("client-uri", "awsdynamodb://?local=true", "A valid aaronland/go-aws-dynamodb URI")

	flag.Parse()

	ctx := context.Background()

	client, err := aa_dynamodb.NewClientWithURI(ctx, *client_uri)

	if err != nil {
		log.Fatalf("Failed to create new client, %v", err)
	}

	input := &aws_dynamodb.ListTablesInput{}

	rsp, err := client.ListTables(ctx, input)

	if err != nil {
		log.Fatalf("Failed to list tables, %v", err)
	}

	for _, t := range rsp.TableNames {
		fmt.Println(t)
	}
}
