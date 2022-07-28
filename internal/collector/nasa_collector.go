package collector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/date_range"
)

const (
	nasaUrl = "https://api.nasa.gov/planetary/apod"
)

type nasaImageCollector struct {
	apiKey string

	imgFetcher imageFetcher
}

func NewNasaImageCollector(semaphore chan struct{}, apiKey string) ImageCollector {
	return &nasaImageCollector{
		apiKey: apiKey,
		imgFetcher: nasaImageFetcher{
			apiKey:    apiKey,
			semaphore: semaphore,
		},
	}
}

type nasaResponse struct {
	Url string `json:"url"`
}

func (imgCollector nasaImageCollector) GetImages(from, to time.Time) ([]string, error) {
	results := []string{}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	dates := date_range.GetDateRangeSlice(from, to)

	fetchChan := make(chan struct {
		url string
		err error
	}, len(dates))

	for _, date := range dates {
		go func(date string) {
			url, err := imgCollector.imgFetcher.fetchImage(date)
			fetchChan <- struct {
				url string
				err error
			}{url, err}
		}(date.Format("2006-01-02"))
	}

	for i := 0; i < len(dates); i++ {
		fetchResult := <-fetchChan
		if fetchResult.err != nil {
			return results, fetchResult.err
		}
		results = append(results, fetchResult.url)
	}

	return results, nil
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
