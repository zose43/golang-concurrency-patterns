package main

import (
	"fmt"
	"log"
	"time"
)

var list *LinkedList

func main() {
	start := time.Now()
	if list == nil {
		log.Fatal("not initialized list")
	}

	streamCh := make(chan *Stream, 10)
	done := make(chan struct{})

	<-done
	fmt.Printf("process took %s\n", time.Since(start))
}

func init() {
	// todo parse csv and fill stream linked list
}

func producer() {

}

func consumer() {

}
