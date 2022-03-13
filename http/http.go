package http

import (
	"net/http"
	"photosplatter/assets"
	"photosplatter/gallery"
)

type photoSet struct {
	Photos []gallery.Photo
}

func Index(w http.ResponseWriter, _ *http.Request) {
	err := assets.Index.Execute(w, photoSet{Photos: gallery.GetGallery().Photos})
	if err != nil {
		panic(err)
	}
}
