package http

import (
	"fmt"
	"net/http"
	"photosplatter/assets"
	"photosplatter/gallery"
)

type photoSet struct {
	Photos []gallery.Photo
}

func Index(w http.ResponseWriter, _ *http.Request) {
	err := assets.Index.Execute(w, []byte{})
	if err != nil {
		panic(err)
	}
}

func AllPhotos(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("content-type", "application/json")
	gal := gallery.GetGallery()
	_, err := w.Write([]byte{'['})
	if err != nil {
		panic(err)
	}
	n := len(gal.Photos)
	for i := 0; i < n; i++ {
		_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, gal.Photos[i].Name)))
		if i == n-1 {
			break
		}
		_, _ = w.Write([]byte{','})
	}
	_, err = w.Write([]byte{']'})
	if err != nil {
		panic(err)
	}
}
