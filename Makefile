GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

dynamodb-local:
	docker run --rm -it -p 8000:8000 amazon/dynamodb-local

cli-auth:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-sts-session cmd/aws-sts-session/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-mfa-session cmd/aws-mfa-session/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-get-credentials cmd/aws-get-credentials/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-cognito-credentials cmd/aws-cognito-credentials/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-set-env cmd/aws-set-env/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-sign-request cmd/aws-sign-request/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-credentials-json-to-ini cmd/aws-credentials-json-to-ini/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/aws-imds-credentials cmd/aws-imds-credentials/main.go

cli-cloudfront:
	go build -mod vendor -o bin/invalidate cmd/invalidate/main.go

cli-cloudwatch:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/log-groups cmd/log-groups/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/log-group-streams cmd/log-group-streams/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/log-stream-events cmd/log-stream-events/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/empty-streams cmd/empty-streams/main.go

cli-ecs:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ecs-launch-task cmd/ecs-launch-task/main.go

cli-lambda:
	go build -mod $(GOMOD) -ldflags "$(LDFLAGS)" -o bin/invoke cmd/invoke/main.go
	go build -mod $(GOMOD) -ldflags "$(LDFLAGS)" -o bin/functionurl cmd/functionurl/main.go
