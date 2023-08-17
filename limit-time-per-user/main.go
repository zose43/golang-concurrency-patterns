package main

import (
	"log"
	"math/rand"
	"time"
)

type User struct {
	ID          int
	TimeAllowed int
	IsPremium   bool
}

func main() {
	u := &User{ID: 1, IsPremium: false}
	handleWithTimeout(u)
}

func handleWithTimeout(u *User) bool {
	done := make(chan struct{})
	go func() {
		process(done)
	}()

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-done:
			log.Print("done")
			return true
		case <-ticker.C:
			u.TimeAllowed++
			if u.TimeAllowed > 4 && !u.IsPremium {
				log.Print("free time is over")
				return false
			}
		}
	}
}

func process(done chan<- struct{}) {
	// some processing
	rand.NewSource(time.Now().UnixNano())
	n := rand.Intn(8)
	time.Sleep(time.Duration(n) * time.Second)
	done <- struct{}{}
}
