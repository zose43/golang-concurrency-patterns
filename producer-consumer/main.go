package main

import (
	"encoding/csv"
	"fmt"
	"golang-concurrency-patterns/lib/storage"
	"log"
	"time"
)

var list *LinkedList

func main() {
	start := time.Now()
	if list == nil {
		log.Fatal("not initialized list")
	}

	//streamCh := make(chan *Stream, 10)
	done := make(chan struct{})

	<-done
	fmt.Printf("process took %s\n", time.Since(start))
}

func init() {
	csvF, err := storage.OpenFile("airport_data.csv")
	if err != nil {
		log.Fatal("can't load dataset")
	}
	defer func() { _ = csvF.Close() }()

	reader := csv.NewReader(csvF)
	items, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("can't read file %s", err)
	}

	list = &LinkedList{nil, 0}
	for _, el := range items {
		if "#" != el[0] {
			dt, err := time.Parse("2006-01-02 15:04:05 -07:00", el[4])
			if err != nil {
				log.Printf("can't parse date %s", err)
				continue
			}
			stream := &Stream{
				airport:  el[1],
				city:     el[2],
				timezone: el[3],
				date:     dt,
			}
			list.addToHead(stream)
		}
	}
}

func producer() {

}

func consumer() {

}
