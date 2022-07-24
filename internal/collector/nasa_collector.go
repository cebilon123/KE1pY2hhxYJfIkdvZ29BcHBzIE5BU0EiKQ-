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
	nasaUrl        = "https://api.nasa.gov/planetary/apod"
	timeoutSeconds = 10
)

func NewNasaImageCollector(semaphore chan struct{}, apiKey string) ImageCollector {
	return &nasaImageCollector{
		sema:   semaphore,
		apiKey: apiKey,
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

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		for url := range urlsChan {
			results = append(results, url)
		}
	}()

	imgFetcher := &imageFetcher{
		apiKey:    imgCollector.apiKey,
		semaphore: imgCollector.sema,
		wg:        &wg,
	}

	for dateIterate := iterate.DateRange(from, to); ; {
		date, isNext := dateIterate()

		if !isNext || date.IsZero() {
			break
		}

		wg.Add(1)
		go imgFetcher.handleNasaFetch(date.Format("2006-01-02"), urlsChan, errChan)
	}

	wg.Wait()
	close(errChan)

	err := <-errChan

	return results, err
}

// imageFetcher is used to fetch images
// from nasa api.
type imageFetcher struct {
	// key to api
	apiKey string

	// semaphore used maintain max concurrent requests
	semaphore chan struct{}
	wg        *sync.WaitGroup
}

// handleNasaFetch handles fetch to nasa api.
// formatedDate is date in format "2006-01-02"
func (imgFetcher *imageFetcher) handleNasaFetch(formatedDate string, urlsChan chan<- string, errChan chan<- error) {
	imgFetcher.semaphore <- struct{}{}

	defer func() {
		if err := recover(); err != nil {
			log.Println("NASA api fetch panicked")
		}
		imgFetcher.wg.Done()
		<-imgFetcher.semaphore
	}()

	url := fmt.Sprintf("%s?api_key=%s&date=%s", nasaUrl, "y5cnFTkqJzcsSp0I9lAfaaRN6ZpahfIrSswujolO", formatedDate)

	response, err := http.Get(url)
	if err != nil {
		errChan <- err
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		errChan <- err
		return
	}

	if response.StatusCode >= 400 {
		errChan <- errors.New(string(responseData))
		return
	}

	var nasaRes nasaResponse
	json.Unmarshal(responseData, &nasaRes)

	if len(nasaRes.Url) > 0 {
		urlsChan <- nasaRes.Url
	}
}
