package main

import (
	"encoding/json"
	"fmt"
	"golang-concurrency-patterns/lib/storage"
	"io"
	"log"
	"net/http"
)

var airports = make(map[string][]string)

func main() {
	mustFillData()
	go http.HandleFunc("/airports", handler)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	notFound := func(city, msg string, w http.ResponseWriter) {
		_, _ = fmt.Fprint(w, msg)
	}

	city := r.URL.Query().Get("city")
	if len(city) == 0 {
		notFound(city, "empty param CITY", w)
		return
	}

	l := NewLRUCache(3)
	res := l.Get(city)
	if res == nil {
		res, ok := airports[city]
		if !ok {
			notFound(city, fmt.Sprintf("can't find %s", city), w)
			return
		}
		l.Put(city, res)
		printAirports(res, w)
	} else {
		printAirports(res.([]string), w)
	}
}

func printAirports(airports []string, w io.Writer) {
	for _, airport := range airports {
		_, _ = fmt.Fprintf(w, "%s\n", airport)
	}
}

func mustFillData() {
	bt, err := storage.OpenData("city_airports.json")
	if err != nil {
		log.Fatal(err)
	}

	var collection []map[string]string
	if err := json.Unmarshal(bt, &collection); err != nil {
		log.Fatal(err)
	}

	for _, airport := range collection {
		city, ok := airport["city"]
		if !ok {
			log.Printf("can't match city %v", airport)
			continue
		}
		a, ok := airport["airport"]
		if !ok {
			log.Printf("can't match airport %v", airport)
			continue
		}
		airports[city] = append(airports[city], a)
	}
}
