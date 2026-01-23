// aws-get-credentials is a command line tool to emit one or more keys from
// a given profile in an AWS .credentials file.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/auth"
	"gopkg.in/ini.v1"
)

func main() {

	profile := flag.String("profile", "default", "A valid AWS credentials profile")

	flag.Parse()

	creds, err := auth.NewCredentials()

	if err != nil {
		log.Fatalf("Failed to create new config, %v", err)
	}

	ini_config, err := ini.Load(creds.Path)

	if err != nil {
		log.Fatalf("Failed to load credentials file, %v", err)
	}

	sect := ini_config.Section(*profile)

	for _, cred := range flag.Args() {

		k, err := sect.GetKey(cred)

		if err != nil {
			log.Fatalf("Failed to get config key %s, %v", cred, err)
		}

		fmt.Println(k.Value())
	}

	os.Exit(0)
}
