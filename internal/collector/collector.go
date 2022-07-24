package collector

import "time"

// ImageCollector enables functionality to
// fetch images from image provider
type ImageCollector interface {
	// Get image url from provider by passing
	// from and to dates. Returns url and nil error
	// if fetched successfully
	GetImages(from, to time.Time) ([]string, error)
}

type nasaImageCollector struct {
	sema   chan struct{}
	apiKey string
}
