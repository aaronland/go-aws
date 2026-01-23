// aws-set-env is a command line tool to assign required AWS authentication environment
// variables for a given profile in a AWS .credentials file.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/aaronland/go-aws/v3/auth"
	"gopkg.in/ini.v1"
)

func main() {

	profile := flag.String("profile", "default", "A valid AWS credentials profile")
	token := flag.Bool("session-token", true, "Require AWS_SESSION_TOKEN environment variable")

	flag.Parse()

	creds, err := auth.NewCredentials()

	if err != nil {
		log.Fatalf("Failed to create new credentials, %v", err)
	}

	ini_config, err := ini.Load(creds.Path)

	if err != nil {
		log.Fatalf("Failed to load config file, %v", err)
	}

	sect := ini_config.Section(*profile)

	aws_creds := map[string]string{
		"aws_access_key_id":     "AWS_ACCESS_KEY_ID",
		"aws_secret_access_key": "AWS_SECRET_ACCESS_KEY",
	}

	if *token {
		aws_creds["aws_session_token"] = "AWS_SESSION_TOKEN"
	}

	for ini_prop, env_var := range aws_creds {

		k, err := sect.GetKey(ini_prop)

		if err != nil {
			log.Fatalf("Failed to load config key %s, %v", ini_prop, err)
		}

		err = os.Setenv(env_var, k.Value())

		if err != nil {
			log.Fatalf("Failed to set environment variable %s, %v", env_var, err)
		}
	}

}
