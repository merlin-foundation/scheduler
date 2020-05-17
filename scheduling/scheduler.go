package scheduling

import (
	"context"
	"sync"
	"time"
)

//Scheduler model
type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

//Job type function
type Job func(ctx context.Context)

//NewScheduler creates a scheduler
func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

//Add scheduler
func (s *Scheduler) Add(ctx context.Context, j Job, interval time.Duration) {
	ctx, cancel := context.WithCancel(ctx)

	s.cancellations = append(s.cancellations, cancel)
	s.wg.Add(1)
	go s.proccess(ctx, j, interval)
}

//Stop jobs
func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) proccess(ctx context.Context, j Job, interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			return
		}
	}
}
