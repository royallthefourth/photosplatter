package main

import (
	"errors"
	"github.com/h2non/filetype"
	"io/ioutil"
	"log"
	"os"
)

func readPath(p string) ([]os.DirEntry, error) {
	if p == "" {
		return nil, errors.New("p flag must be nonempty")
	}

	err := os.Chdir(p)
	if err != nil {
		return nil, err
	}
	return os.ReadDir(p)
}

func onlyPhotos(files []os.DirEntry) []string {
	images := make([]string, 0)
	for _, file := range files {
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Printf("Error while reading %s: %s", file.Name(), err.Error())
			continue
		}
		if filetype.IsImage(data) {
			images = append(images, file.Name())
		}
	}
	return images
}
