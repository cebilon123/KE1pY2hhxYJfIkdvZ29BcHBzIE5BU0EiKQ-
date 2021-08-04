// Package domain contains whole domain logic of application,
// i.e. models or interfaces.
package domain

import "time"

// ImageProvider provides functions to handle Image fetch.
type ImageProvider interface {
	// GetImagesUrls returns slice which contains at least one image. If there is none,
	// function returns error.
	GetImagesUrls(startDate, endDate time.Time) ([]Image, error)
}
