package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"photosplatter/assets/html"
	"sort"
	"time"
)

var path = flag.String("path", "", "path to the directory containing images")
var port = flag.Int("port", 8080, "http port")

func main() {
	flag.Parse()

	var (
		err    error
		photos []photo
	)

	start := time.Now()
	if photos, err = readPath(*path); err != nil {
		log.Fatalf(err.Error())
	}
	sort.Sort(sort.Reverse(ByDate(photos)))
	finish := time.Now()
	log.Printf("Scanned %d files in %.3fs\n", len(photos), float32(finish.Sub(start).Milliseconds())/1000)

	mux := http.NewServeMux()
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(*path))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(html.Assets))))
	mux.HandleFunc("/", index(photos))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), mux))
}

type photo struct {
	Created time.Time
	Name    string
}
