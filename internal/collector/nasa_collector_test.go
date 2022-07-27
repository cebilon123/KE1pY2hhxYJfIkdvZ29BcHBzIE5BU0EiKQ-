package collector

import (
	"errors"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

var (
	imagesUrls = []string{
		"url1",
		"url2",
		"url3",
	}
	validFirstDate  = time.Date(2001, 1, 1, 1, 1, 1, 1, time.UTC)
	validSecondDate = time.Date(2001, 1, 2, 1, 1, 1, 1, time.UTC)
	validThirdDate  = time.Date(2001, 1, 3, 1, 1, 1, 1, time.UTC)
)

type imageFetcherMock struct {
}

func (m *imageFetcherMock) fetchImage(date string) (string, error) {

	// mock some kind of network latency
	time.Sleep(125 * time.Millisecond)

	if date == validFirstDate.Format("2006-01-02") {
		return imagesUrls[0], nil
	}
	if date == validSecondDate.Format("2006-01-02") {
		return imagesUrls[1], nil
	}
	if date == validThirdDate.Format("2006-01-02") {
		return imagesUrls[2], nil
	}

	return "", errors.New("example error while fetching")
}

func TestCollector_FetchImages(t *testing.T) {
	semaphore := make(chan struct{}, 3)
	imgCollector := &nasaImageCollector{
		sema:       semaphore,
		apiKey:     "test",
		imgFetcher: &imageFetcherMock{},
	}

	type input struct {
		from time.Time
		to   time.Time
	}

	type expected struct {
		images []string
		err    error
	}

	testCases := []struct {
		input    input
		expected expected
	}{
		{
			input: input{
				from: validFirstDate,
				to:   validThirdDate,
			},
			expected: expected{
				images: imagesUrls,
				err:    nil,
			},
		},
	}

	for _, tt := range testCases {
		actualUrls, err := imgCollector.GetImages(tt.input.from, tt.input.to)
		areElementsInArray := func(urls []string, expectedUrls []string) bool {
			for _, url := range expectedUrls {
				if !slices.Contains(urls, url) {
					return false
				}
			}

			return true
		}

		if len(actualUrls) != len(tt.expected.images) || err != tt.expected.err || !areElementsInArray(actualUrls, tt.expected.images) {
			t.Errorf("GetImages: expected: %v, actual: %v", tt.expected.images, actualUrls)
		}
	}
}
