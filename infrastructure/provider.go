package infrastructure

import (
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/domain"
	"net/url"
	"time"
)

const nasaBaseUrl = "https://api.nasa.gov/planetary/apod"

// NasaImageProvider represents nasa image provider which implements
// ImageProvider interface.
type NasaImageProvider struct {
	apiUrl url.URL
}

// NewNasaImageProvider creates new instance of NasaImageProvider.
func NewNasaImageProvider() *NasaImageProvider {
	return &NasaImageProvider{apiUrl: url.URL{
		Path: nasaBaseUrl,
	}}
}


func (n NasaImageProvider) GetImagesUrls(startDate, endDate time.Time) ([]domain.Image, error) {
	panic("implement me")
}

