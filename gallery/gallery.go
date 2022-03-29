package gallery

import (
	"errors"
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

type Gallery interface {
	Photos() []Photo
	ScanForChanges()
}

func NewDiskBacked(path string) (DiskBacked, error) {
	if path == "" {
		return DiskBacked{}, errors.New("path must be nonempty")
	}

	return DiskBacked{
		path: path,
	}, os.Chdir(path)
}

type DiskBacked struct {
	mu       sync.Mutex
	path     string
	photos   []Photo
	rawCount int
}

func (d *DiskBacked) Photos() []Photo {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.photos
}

func (d *DiskBacked) ScanForChanges() {
	var (
		start, finish time.Time
	)
	for {
		rawFiles, err := os.ReadDir(d.path)
		if err != nil {
			log.Printf("Dir scan error: %s", err.Error())
		}
		newCount := len(rawFiles)

		if d.rawCount != newCount {
			d.mu.Lock()
			start = time.Now()
			d.rawCount = len(rawFiles)
			d.photos = scanFiles(rawFiles)
			finish = time.Now()
			log.Printf("Scanned %d photos in %.3f seconds", d.rawCount, float32(finish.Sub(start).Milliseconds())/1000)
			d.mu.Unlock()
		}

		time.Sleep(time.Minute)
	}
}

func scanFiles(rawFiles []os.DirEntry) []Photo {
	photos := onlyPhotos(rawFiles)
	sort.Sort(sort.Reverse(ByDate(photos)))
	return photos
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
