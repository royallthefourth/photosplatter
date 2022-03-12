package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"photosplatter/assets/html"
)

var path = flag.String("path", "", "path to the directory containing images")
var port = flag.Int("port", 8080, "http port")

func main() {
	flag.Parse()

	var (
		err      error
		rawFiles []os.DirEntry
	)

	if rawFiles, err = readPath(*path); err != nil {
		log.Fatalf(err.Error())
	}

	mux := http.NewServeMux()
	mux.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(*path))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(html.Assets))))
	mux.HandleFunc("/", index(onlyPhotos(rawFiles)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), mux))
}
