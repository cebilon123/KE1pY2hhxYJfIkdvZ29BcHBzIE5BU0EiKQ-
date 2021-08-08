package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/customError"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/domain"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/server"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
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

// NasaImageProvider represents nasa image provider which implements
// ImageProvider interface.
type NasaImageProvider struct {
	apiUrl url.URL
	err    error
}

// requestHandler handles request to nasa api
type requestHandler struct {
	id int
}

var availableRequestHandlers chan requestHandler

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

	return &NasaImageProvider{apiUrl: apiUrl}
}

func (n *NasaImageProvider) GetImagesUrls(startDate, endDate time.Time) ([]domain.Image, error) {
	u := getUrlWithDates(startDate, endDate, n.apiUrl)
	h := <-availableRequestHandlers
	res, err := h.getResponse(u.String())
	n.err = err
	ret := n.UnmarshalBody(n.ReadBody(res))

	availableRequestHandlers <- h

	return ret, n.err
}

func (r *requestHandler) getResponse(url string) (*http.Response, error) {
	return http.Get(url)
}

func (n *NasaImageProvider) ReadBody(r *http.Response) []byte{
	if r.StatusCode != http.StatusOK {
		n.err = customError.Server{Message: "Nasa request returned status other than 200"}
		return nil
	}
	if n.err != nil {
		return nil
	}

	body, err := ioutil.ReadAll(r.Body)
	n.err = err

	return body
}

func (n *NasaImageProvider) UnmarshalBody(b []byte) []domain.Image {
	if n.err != nil {
		return nil
	}

	var ret []domain.Image
	if err := json.Unmarshal(b, &ret); err != nil {
		n.err = err
		return nil
	}

	return ret
}

func getUrlWithDates(startDate, endDate time.Time, ur url.URL) url.URL {
	u := ur
	q := u.Query()
	q.Set(startDateParam, fmt.Sprintf("%d-%02d-%02d", startDate.Year(), startDate.Month(), startDate.Day()))
	q.Set(endDateParam, fmt.Sprintf("%d-%02d-%02d", endDate.Year(), endDate.Month(), endDate.Day()))
	u.RawQuery = q.Encode()

	return u
}

func init() {
	maxRq, err := strconv.Atoi(os.Getenv(server.MaxConcurrentApiCalls))
	if err != nil {
		maxRq = 5
	}

	availableRequestHandlers = make(chan requestHandler, maxRq)
	for i := 0; i < maxRq; i++ {
		availableRequestHandlers <- requestHandler{id: i}
	}
}
