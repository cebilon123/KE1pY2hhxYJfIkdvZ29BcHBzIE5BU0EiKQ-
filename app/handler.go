package app

import (
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/domain"
	"net/http"
	"time"
)

const (
	startDateParam = "start_date"
	endDateParam = "end_date"
)

// NasaPicturesHandler handles image url request
func NasaPicturesHandler(w http.ResponseWriter, r *http.Request, ip domain.ImageProvider) {
	ctx := r.Context()
	start,err := time.Parse("2006-01-02","2021-08-01")
	end,err := time.Parse("2006-01-02","2021-08-04")
	if err != nil {

	}
	ip.GetImagesUrls(start,end)
	ctx.Done()
}
