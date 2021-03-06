package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/merlin-foundation/scheduler/scheduling"
)

func main() {
	ctx := context.Background()

	worker := scheduling.NewScheduler()
	worker.Add(ctx, parseSubscriptionData, time.Second*5)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}

func parseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 5)
	fmt.Printf("testando %s\n", time.Now().String())
}
