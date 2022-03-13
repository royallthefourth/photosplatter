package gallery

import (
	"github.com/h2non/filetype"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
	"time"
)

type Photo struct {
	Created time.Time
	Name    string
}

type Gallery struct {
	Photos   []Photo
	RawCount int
}

var (
	gal Gallery
	mu  sync.Mutex
)

func GetGallery() Gallery {
	mu.Lock()
	defer mu.Unlock()
	return gal
}

func WatchFiles(p string) {
	var (
		start, finish time.Time
	)
	for {
		rawFiles, err := os.ReadDir(p)
		if err != nil {
			log.Printf("Dir scan error: %s", err.Error())
		}
		newCount := len(rawFiles)

		if gal.RawCount != newCount {
			mu.Lock()
			start = time.Now()
			gal = scanFiles(rawFiles)
			finish = time.Now()
			log.Printf("Scanned %d photos in %.3f seconds", gal.RawCount, float32(finish.Sub(start).Milliseconds())/1000)
			mu.Unlock()
		}

		time.Sleep(time.Minute)
	}
}

func scanFiles(rawFiles []os.DirEntry) Gallery {
	newGal := Gallery{
		Photos:   onlyPhotos(rawFiles),
		RawCount: len(rawFiles),
	}
	sort.Sort(sort.Reverse(ByDate(newGal.Photos)))
	return newGal
}

func onlyPhotos(files []os.DirEntry) []Photo {
	photos := make([]Photo, 0)
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
			photos = append(photos, Photo{
				Created: creationTime,
				Name:    file.Name(),
			})
		}
	}
	return photos
}
