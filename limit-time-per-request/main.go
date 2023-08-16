package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	h := getHost()
	user := &User{ID: 1, IsPremium: false}
	serv := NewService(5*time.Second, h)
	serv.User = user

	ctx, cancel := context.WithTimeout(context.TODO(), 4*time.Second)
	defer cancel()

	res := timeoutProcess(ctx, serv)
	if res {
		fmt.Println("done")
	}
}

func getHost() string {
	var host string
	flag.StringVar(&host, "h", "https://httpbin.org", "rest service")
	flag.Parse()
	return host
}

func timeoutProcess(ctx context.Context, serv *Service) bool {
	done := make(chan error)

	go func() {
		done <- serv.Process("json")
	}()

	for {
		select {
		case err := <-done:
			if err != nil {
				log.Print(err)
				return false
			}
			return true
		case <-ctx.Done():
			if !serv.User.IsPremium {
				log.Print("timeout")
				return false
			}
		}
	}
}
