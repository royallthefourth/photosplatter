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

func AllPhotos(gal gallery.Gallery) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("content-type", "application/json")
		_, err := w.Write([]byte{'['})
		if err != nil {
			panic(err)
		}
		photos := gal.Photos()
		n := len(photos)
		for i := 0; i < n; i++ {
			_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, photos[i].Name)))
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
}
