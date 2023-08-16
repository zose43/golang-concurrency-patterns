package main

import (
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

	res := timeoutProcess(serv)
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

func timeoutProcess(serv *Service) bool {
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
		case <-time.After(serv.AllowedDuration):
			if !serv.User.IsPremium {
				log.Print("timeout")
				return false
			}
		}
	}
}
