package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/domain"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/server"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	nasaScheme     = "https"
	nasaHost       = "api.nasa.gov"
	nasaPath       = "planetary/apod"
	startDateParam = "start_date"
	endDateParam   = "end_date"
	apiKeyParam    = "api_key"
)

var (
	instance     *NasaImageProvider
	once         sync.Once
	currRequests int8
)

// NasaImageProvider represents nasa image provider which implements
// ImageProvider interface.
type NasaImageProvider struct {
	apiUrl        url.URL
	maxApiRequest int8
	mu            sync.Mutex
}

// GetNasaImageProvider returns nasa image provider instance.
func GetNasaImageProvider() *NasaImageProvider {
	once.Do(func() {
		instance = NewNasaImageProvider()
	})

	return instance
}

// NewNasaImageProvider creates new instance of NasaImageProvider.
func NewNasaImageProvider() *NasaImageProvider {
	apiUrl := url.URL{
		Host:   nasaHost,
		Scheme: nasaScheme,
		Path:   nasaPath,
	}

	q := apiUrl.Query()
	q.Set(apiKeyParam, os.Getenv(server.ApiKey))
	apiUrl.RawQuery = q.Encode()

	maxRequestsNum, err := strconv.Atoi(os.Getenv(server.MaxConcurrentApiCalls))
	if err != nil {
		maxRequestsNum = 5
	}

	return &NasaImageProvider{apiUrl: apiUrl, maxApiRequest: int8(maxRequestsNum)}
}

func (n *NasaImageProvider) GetImagesUrls(startDate, endDate time.Time) ([]domain.Image, error) {
	n.mu.Lock()
	defer func() {
		currRequests--
		n.mu.Unlock()
	}()

	for {
		if currRequests < n.maxApiRequest {
			break
		}
	}

	currRequests++
	fmt.Printf("Curr requests: %v\n", currRequests)

	u := getUrlWithDates(startDate, endDate, n.apiUrl)
	res, err := http.Get(u.String())
	fmt.Printf("Calling: %s\n", u.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ret := make([]domain.Image, 1)
	if err := json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func getUrlWithDates(startDate, endDate time.Time, ur url.URL) url.URL {
	u := ur
	q := u.Query()
	q.Set(startDateParam, fmt.Sprintf("%d-%02d-%02d", startDate.Year(), startDate.Month(), startDate.Day()))
	q.Set(endDateParam, fmt.Sprintf("%d-%02d-%02d", endDate.Year(), endDate.Month(), endDate.Day()))
	u.RawQuery = q.Encode()

	return u
}
