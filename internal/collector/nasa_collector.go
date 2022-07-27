package collector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/iterate"
)

const (
	nasaUrl = "https://api.nasa.gov/planetary/apod"
)

type nasaImageCollector struct {
	sema   chan struct{}
	apiKey string

	imgFetcher imageFetcher
}

func NewNasaImageCollector(semaphore chan struct{}, apiKey string) ImageCollector {
	return &nasaImageCollector{
		sema:   semaphore,
		apiKey: apiKey,
		imgFetcher: nasaImageFetcher{
			apiKey:    "y5cnFTkqJzcsSp0I9lAfaaRN6ZpahfIrSswujolO",
			semaphore: semaphore,
		},
	}
}

type nasaResponse struct {
	Url string `json:"url"`
}

func (imgCollector nasaImageCollector) GetImages(from, to time.Time) ([]string, error) {
	results := []string{}
	urlsChan := make(chan string)
	errChan := make(chan error, 1)

	var wg sync.WaitGroup
	var fetchError error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	for dateIterate := iterate.DateRange(from, to); ; {
		date, isNext := dateIterate()

		if !isNext || date.IsZero() {
			break
		}

		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			url, err := imgCollector.imgFetcher.fetchImage(date.Format("2006-01-02"))
			if err != nil {
				errChan <- err
				return
			}
			urlsChan <- url
		}()
	}

	go func() {
		for url := range urlsChan {
			results = append(results, url)
		}
	}()

	go func() {
		fetchError = <-errChan
	}()

	wg.Wait()

	return results, fetchError
}

// imageFetcher is a interface that must be implemented
// by nasaImageFetcher. Used to make dependency inversion
// available
type imageFetcher interface {
	fetchImage(date string) (string, error)
}

// nasaImageFetcher is used to fetch images
// from nasa api.
type nasaImageFetcher struct {
	// key to api
	apiKey string
	// semaphore used maintain max concurrent requests
	semaphore chan struct{}
}

func (imgFetcher nasaImageFetcher) fetchImage(date string) (string, error) {
	imgFetcher.semaphore <- struct{}{}

	defer func() {
		if err := recover(); err != nil {
			log.Println("NASA api fetch panicked")
		}
		<-imgFetcher.semaphore
	}()

	url := fmt.Sprintf("%s?api_key=%s&date=%s", nasaUrl, imgFetcher.apiKey, date)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	if response.StatusCode >= 400 {
		return "", errors.New(string(responseData))
	}

	var nasaRes nasaResponse
	json.Unmarshal(responseData, &nasaRes)

	return nasaRes.Url, nil
}
