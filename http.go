package main

import (
	"net/http"
	"photosplatter/assets/html"
)

type photoSet struct {
	Photos []photo
}

func index(photos []photo) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		err := html.Index.Execute(w, photoSet{Photos: photos})
		if err != nil {
			panic(err)
		}
	}
}
