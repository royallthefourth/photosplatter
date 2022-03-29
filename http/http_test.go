package http

import (
	"encoding/json"
	"github.com/icrowley/fake"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"photosplatter/gallery"
	"sort"
	"testing"
	"time"
)

func fakeGallery(n int) gallery.InMemory {
	photos := make([]gallery.Photo, n)
	for i := 0; i < n; i++ {
		photos[i] = gallery.Photo{
			Created: time.Date(
				fake.Year(1900, time.Now().Year()),
				time.Month(fake.MonthNum()),
				fake.Day(),
				rand.Intn(24),
				rand.Intn(60),
				rand.Intn(60),
				0,
				time.UTC),
			Name: fake.ProductName(),
		}
	}
	return gallery.InMemory{P: photos}
}

func TestAllPhotos(t *testing.T) {
	handler := AllPhotos(fakeGallery(5))
	req := httptest.NewRequest(http.MethodGet, "http://localhost/api/photos", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if !json.Valid(body) {
		t.Errorf("Expected valid json but got %s", string(body))
	}
}

func BenchmarkAllPhotos(b *testing.B) {
	gal := fakeGallery(1_000_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		handler := AllPhotos(gal)
		req := httptest.NewRequest(http.MethodGet, "http://localhost/api/photos", nil)
		w := httptest.NewRecorder()
		handler(w, req)
	}
}

func BenchmarkSort(b *testing.B) {
	g := fakeGallery(1_000_000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(sort.Reverse(gallery.ByDate(g.Photos())))
	}
}
