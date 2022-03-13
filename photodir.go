package main

import (
	"errors"
	"github.com/h2non/filetype"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func readPath(p string) ([]photo, error) { // TODO return PhotoTree
	if p == "" {
		return nil, errors.New("p flag must be nonempty")
	}

	err := os.Chdir(p)
	if err != nil {
		return nil, err
	}
	rawFiles, err := os.ReadDir(p)
	if err != nil {
		return nil, err
	}
	return onlyPhotos(rawFiles), err
}

func onlyPhotos(files []os.DirEntry) []photo { // TODO return PhotoTree
	photos := make([]photo, 0)
	var creationTime time.Time
	for _, file := range files {
		data, err := ioutil.ReadFile(file.Name())
		if err != nil {
			log.Printf("Error while reading %s: %s", file.Name(), err.Error())
			continue
		}

		creationTime, err = getCreationTime(file)
		if err != nil {
			log.Printf("Error getting creation time %s: %s", file.Name(), err.Error())
			continue
		}

		if filetype.IsImage(data) {
			photos = append(photos, photo{
				Created: creationTime,
				Name:    file.Name(),
			})
		}
	}
	return photos
}
