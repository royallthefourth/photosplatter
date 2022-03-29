package http

import (
	"encoding/json"
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

func AllPhotos(gal gallery.Gallery) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("content-type", "application/json")
		photos := gal.Photos()
		n := len(photos)
		out := make([]string, n)
		for i := 0; i < n; i++ {
			out[i] = photos[i].Name
		}
		b, err := json.Marshal(out)
		if err != nil {
			panic(err)
		}
		_, err = w.Write(b)
		if err != nil {
			panic(err)
		}
	}
}
