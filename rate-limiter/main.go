package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"time"
)

const Depth = 2

func main() {
	if len(os.Args) < 2 {
		log.Fatal("link is empty")
	}
	link := os.Args[1]
	if !validateLink(link) {
		log.Fatalf("%q is not a link", link)
	}

	ticker := time.NewTicker(500 * time.Millisecond)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go process(ticker.C, link, Depth, &wg)
	wg.Wait()

	log.Print("done")
}

func validateLink(link string) bool {
	res, _ := regexp.MatchString(`https?`, link)
	return res
}

func process(ticker <-chan time.Time, link string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth < 0 {
		return
	}

	<-ticker

	title, links, err := Crawler(link)
	if err != nil {
		//todo check consistent errors
		log.Printf("can't get title on page %q %s\n", link, err)
	}
	fmt.Printf("%s --> %q\n", link, title)

	wg.Add(len(links))
	for i := range links {
		go process(ticker, links[i], depth-1, wg)
	}
}
