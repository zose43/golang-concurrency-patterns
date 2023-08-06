package main

import (
	"context"
	"log"
	"sync"
)

type Wp struct {
	workerNum int
	jobs      chan Job
	results   chan Result
	done      chan struct{}
}

func worker(ctx context.Context, wg *sync.WaitGroup, input <-chan Job, output chan<- Result) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-input:
			if !ok {
				return
			}
			output <- job.execute(ctx)
		case <-ctx.Done():
			log.Printf("cancell %v", ctx.Err())
			output <- Result{Err: ctx.Err()}
		}
	}
}

func (wp Wp) Run(ctx context.Context) {
	wg := sync.WaitGroup{}
	for i := 0; i < wp.workerNum; i++ {
		wg.Add(1)
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.done)
	close(wp.results)
}

func (wp Wp) Results() <-chan Result {
	return wp.results
}

func (wp Wp) GenerateFrom(jobsColl []Job) {
	for i := range jobsColl {
		wp.jobs <- jobsColl[i]
	}
	close(wp.jobs)
}

func NewWp(workerNum int) Wp {
	return Wp{
		workerNum: workerNum,
		jobs:      make(chan Job),
		results:   make(chan Result),
		done:      make(chan struct{}),
	}
}
