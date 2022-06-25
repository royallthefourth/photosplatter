package main

import (
	"flag"
	"fmt"
	"log"
	libhttp "net/http"
	"photosplatter/assets"
	"photosplatter/gallery"
	"photosplatter/http"
)

var path = flag.String("path", "", "path to the directory containing images")
var port = flag.Int("port", 8080, "http port")

func main() {
	flag.Parse()

	gal, err := gallery.NewDiskBacked(*path)
	if err != nil {
		log.Fatalf(err.Error())
	}

	go gal.ScanForChanges()

	mux := libhttp.NewServeMux()
	mux.Handle("/images/", libhttp.StripPrefix("/images/", libhttp.FileServer(libhttp.Dir(*path))))
	mux.Handle("/assets/", libhttp.StripPrefix("/assets/", libhttp.FileServer(libhttp.FS(assets.Assets))))

	mux.HandleFunc("/", http.Index)
	mux.HandleFunc("/api/photos", http.AllPhotos(&gal))
	fmt.Printf("Listening on http://0.0.0.0:%d\n", *port)
	log.Fatal(libhttp.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), mux))
}
