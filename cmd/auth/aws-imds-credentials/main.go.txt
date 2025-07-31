package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/auth"
)

func main() {

	flag.Parse()
	ctx := context.Background()

	creds, err := auth.EC2RoleCredentials(ctx)

	if err != nil {
		log.Fatalf("Failed to retrive credentials, %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	err = enc.Encode(creds)

	if err != nil {
		log.Fatalf("Failed to encode credentials, %v", err)
	}
}
