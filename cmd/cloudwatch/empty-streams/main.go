// empty-streams will list all the CloudWatch log streams with 0 stored bytes and optionally
// remove them.
package main

import (
	"context"
	"flag"
	_ "fmt"
	"log"
	"log/slog"
	"sync"
	"time"

	"github.com/aaronland/go-aws/v3/cloudwatch/logs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func main() {

	var verbose bool

	cloudwatch_uri := flag.String("cloudwatch-uri", "", "...")

	prune := flag.Bool("prune", false, "Remove log streams with no events.")
	dryrun := flag.Bool("dryrun", false, "Go through the motions but don't actually remove any log streams.")

	max_workers := flag.Int("max-workers", 100, "The maximum number of concurrent workers.")

	flag.BoolVar(&verbose, "verbose", false, "")

	flag.Parse()

	if verbose {
		slog.SetLogLoggerLevel(slog.LevelDebug)
		slog.Debug("Verbose logging enabled")
	}

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	throttle := make(chan bool, *max_workers)

	for i := 0; i < *max_workers; i++ {
		throttle <- true
	}

	limiter := time.Tick(200 * time.Millisecond)

	cloudwatch_cl, err := logs.NewClient(ctx, *cloudwatch_uri)

	if err != nil {
		log.Fatalf("Failed to create service, %v", err)
	}

	stream_filter := func(ctx context.Context, s *types.LogStream) (bool, error) {

		if *s.StoredBytes == 0 {
			return true, nil
		}

		return false, nil
	}

	wg := new(sync.WaitGroup)

	for gr, err := range logs.GetLogGroups(ctx, cloudwatch_cl) {

		if err != nil {
			log.Fatalf("Failed to get log groups, %v", err)
		}

		slog.Debug("Process group", "group", *gr.LogGroupName)

		for st, err := range logs.GetLogGroupStreams(ctx, cloudwatch_cl, *gr.LogGroupName, stream_filter) {

			if err != nil {
				log.Fatalf("Failed to get log streams for %s, %v", *gr.LogGroupName, err)
			}

			if *prune {

				slog.Debug("Prune stream", "group", *gr.LogGroupName, "stream", *st.LogStreamName)

				events_opts := &logs.GetLogEventsOptions{
					LogGroupName:  *gr.LogGroupName,
					LogStreamName: *st.LogStreamName,
				}

				has_events := false

				for _, err := range logs.GetLogEvents(ctx, cloudwatch_cl, events_opts) {

					if err != nil {
						log.Fatalf("Failed to get events for %s (%s), %v", *gr.LogGroupName, *st.LogStreamName, err)
					}

					has_events = true
					break
				}

				if has_events {
					slog.Debug("Stream has events, skipping", "group", *gr.LogGroupName, "stream", *st.LogStreamName)
					continue
				}

				wg.Add(1)

				go func(g *types.LogGroup, s *types.LogStream) {

					defer func() {
						throttle <- true
						wg.Done()
					}()

					<-throttle
					<-limiter

					if *dryrun {
						slog.Info("Prune (dryrun)", "group", *g.LogGroupName, "stream", *s.LogStreamName)
						return
					}

					err := pruneStream(ctx, cloudwatch_cl, g, s)

					if err != nil {
						slog.Info("Failed to prune", "group", *g.LogGroupName, "stream", *s.LogStreamName, "error", err)
					} else {
						slog.Info("Successfully pruned stream", "group", *g.LogGroupName, "stream", *s.LogStreamName)
					}

				}(gr, st)

			}
		}
	}

	wg.Wait()
}

func pruneStream(ctx context.Context, cl *cloudwatchlogs.Client, g *types.LogGroup, s *types.LogStream) error {

	opts := &cloudwatchlogs.DeleteLogStreamInput{
		LogGroupName:  g.LogGroupName,
		LogStreamName: s.LogStreamName,
	}

	_, err := cl.DeleteLogStream(ctx, opts)
	return err
}
