GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

dynamodb-local:
	docker run --platform=linux/amd64 --rm -it -p 8000:8000 amazon/dynamodb-local

cli:
	@make cli-auth
	@make cli-cloudfront
	@make cli-cloudwatch
	@make cli-dynamodb
	@make cli-ecs
	@make cli-ec2
	@make cli-lambda

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
	go build -mod vendor -o bin/cloudfront-invalidate cmd/cloudfront-invalidate/main.go
	go build -mod vendor -o bin/cloudfront-distribution-id cmd/cloudfront-distribution-id/main.go

cli-cloudwatch:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/cloudwatch-log-groups cmd/cloudwatch-log-groups/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/cloudwatch-log-group-streams cmd/cloudwatch-log-group-streams/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/cloudwatch-log-stream-events cmd/cloudwatch-log-stream-events/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/cloudwatch-empty-streams cmd/cloudwatch-empty-streams/main.go

cli-dynamodb:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/dynamodb-list-tables cmd/dynamodb-list-tables/main.go

cli-ecs:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ecs-launch-task cmd/ecs-launch-task/main.go

cli-ec2:
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ec2-public-ip cmd/ec2-public-ip/main.go
	go build -mod $(GOMOD) -ldflags="$(LDFLAGS)" -o bin/ec2-reboot-instances cmd/ec2-reboot-instances/main.go

cli-lambda:
	go build -mod $(GOMOD) -ldflags "$(LDFLAGS)" -o bin/lambda-invoke cmd/lambda-invoke/main.go
	go build -mod $(GOMOD) -ldflags "$(LDFLAGS)" -o bin/lambda-functionurl cmd/lambda-functionurl/main.go
