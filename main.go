package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

type service struct{}

func NewService() service {
	return service{}
}

func (s *service) Start(ctx context.Context, wg *sync.WaitGroup) {
	log.Println("[service] starting")
	wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			log.Println("[service] shutting down.")
			defer wg.Done()
			return
		default:
			log.Println("[service] doing some work…")
			time.Sleep(5 * time.Second)
		}
	}
}

type waitChan struct {
	wg sync.WaitGroup
}

func NewWaitChan() waitChan {
	return waitChan{
		wg: sync.WaitGroup{},
	}
}

func (wc *waitChan) WaitGroup() *sync.WaitGroup {
	return &wc.wg
}

func (wc *waitChan) Done() <-chan bool {
	c := make(chan bool)
	go func(chan<- bool) {
		wc.wg.Wait()
		c <- true
	}(c)
	return c
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	wc := NewWaitChan()
	svc := NewService()

	log.Println("[main] starting services")
	go svc.Start(ctx, wc.WaitGroup())

	select {
	case <-ctx.Done():
		stop()
		waitPeriod := 10 * time.Second
		log.Printf("[main] system waiting up to %s before exit.", waitPeriod)

		select {
		case <-time.After(5 * time.Second):
			log.Println("[main] timed out waiting for cleanup; exiting now.")
			return
		case <-wc.Done():
			log.Println("[main] cleanup done; exiting now")
			return
		}
	}
}
