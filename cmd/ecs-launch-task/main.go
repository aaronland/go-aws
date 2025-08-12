package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/aaronland/go-aws/v3/ecs"
	"github.com/sfomuseum/go-flags/multi"
)

func main() {

	task := flag.String("task", "", "The name (and version) of your ECS task.")
	container := flag.String("container", "", "The name of your ECS container.")
	cluster := flag.String("cluster", "", "The name of your ECS cluster.")
	launch_type := flag.String("launch-type", "", "A valid ECS launch type.")
	platform_version := flag.String("platform-version", "", "A valid ECS platform version.")
	public_ip := flag.String("public-ip", "", "A valid ECS public IP string.")

	wait := flag.Bool("wait", false, "")
	wait_interval := flag.Int("wait-interval", 30, "...in seconds")
	wait_timeout := flag.Int("wait-timeout", 300, "...in minutes")

	var subnets multi.MultiString
	var security_groups multi.MultiString

	flag.Var(&subnets, "subnet", "One or more subnets to run your ECS task in.")
	flag.Var(&security_groups, "security-group", "A valid AWS security group to run your task under.")

	session_uri := flag.String("session-uri", "", "A valid aaronland/go-aws-auth URI.")

	flag.Parse()

	ctx := context.Background()

	opts := &ecs.TaskOptions{
		Task:            *task,
		Container:       *container,
		Cluster:         *cluster,
		LaunchType:      *launch_type,
		PlatformVersion: *platform_version,
		PublicIP:        *public_ip,
		Subnets:         subnets,
		SecurityGroups:  security_groups,
	}

	cmd := flag.Args()

	svc, err := ecs.NewService(*session_uri)

	if err != nil {
		log.Fatalf("Failed to create new service, %v", err)
	}

	rsp, err := ecs.LaunchTask(ctx, svc, opts, cmd...)

	if err != nil {
		log.Fatalf("Failed to launch task, %v", err)
	}

	if *wait {

		interval := time.Duration(*wait_interval) * time.Second
		timeout := time.Duration(*wait_timeout) * time.Second

		wait_opts := &ecs.WaitTasksOptions{
			Cluster:  *cluster,
			TaskArns: rsp.Tasks,
			Interval: interval,
			Timeout:  timeout,
		}

		err := ecs.WaitForTasksToComplete(ctx, svc, wait_opts)

		if err != nil {
			log.Fatalf("Failed to wait for tasks to complete, %v", err)
		}
	}

	fmt.Println(rsp)
}
