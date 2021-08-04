package server

import "net/http"

const (
	startDateParam = "start_date"
	endDateParam = "end_date"
)

// NasaPicturesHandler handles image url request
func NasaPicturesHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx.Done()
}
