package main

import (
	"net/http"
	"photosplatter/assets/html"
)

type photoSet struct {
	Photos []string
}

func index(photos []string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		err := html.Index.Execute(w, photoSet{Photos: photos})
		if err != nil {
			panic(err)
		}
	}
}
