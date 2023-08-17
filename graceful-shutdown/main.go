package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

type Process struct {
	ID   int
	done chan struct{}
}

func (p *Process) stop() {
	p.done <- struct{}{}
}

func (p *Process) run() {
	// some logic
	select {
	case <-time.After(10 * time.Second):
		log.Print("done")
	case <-p.done:
		log.Print("\nexit")
		return
	}
}

func main() {
	proc := &Process{1, make(chan struct{})}

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt)
		<-sig
		signal.Reset()
		proc.stop()
	}()

	proc.run()
}
