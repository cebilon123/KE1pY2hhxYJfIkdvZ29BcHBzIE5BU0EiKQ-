package app

import (
	"encoding/json"
	"fmt"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/customError"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/domain"
	"net/http"
	"time"
)

const (
	startDateParam = "start_date"
	endDateParam   = "end_date"
)

// dateValidationRules are rules used to validate dates in areDatesValid() function
// for sake of simplicity here can be passed arbitrary number of time arguments
// bue first one should be startDate (which is earlier than endDate)
var dateValidationRules = map[string]func(dates ...time.Time) bool{
	"start_date should be earlier than end_date": func(dates ...time.Time) bool {
		return dates[0].Before(dates[1])
	},
	// https://apod.nasa.gov/apod/archivepix.html
	"start_date should be greater or equal to 2015-01-01": func(dates ...time.Time) bool {
		return dates[0].After(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC))
	},
	"start_date should be at least today": func(dates ...time.Time) bool {
		return dates[0].Before(time.Now().UTC())
	},
}

type NasaPicturesHandler struct {
	err error
}

func NewNasaPicturesHandler() *NasaPicturesHandler {
	return &NasaPicturesHandler{}
}

// Handle handles image url request
func (n *NasaPicturesHandler) Handle(w http.ResponseWriter, r *http.Request, ip domain.ImageProvider) {
	ctx := r.Context()
	defer ctx.Done()

	sDate, eDate := n.obtainStringDatesFromRequest(r)
	start, end := n.parseToDates(sDate, eDate)

	if n.err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(customError.NewApiErr(n.err.Error()).ToByteSlice())
		return
	}

	datesValid, msg := areDatesValid(start, end)
	if !datesValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr(msg).ToByteSlice())
		return
	}

	res, err := ip.GetImagesUrls(start, end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr(err.Error()).ToByteSlice())
		return
	}

	mRes, err := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Write(mRes)
}

// areDatesValid Checks if date range is valid. Returns true whenever date is valid
// and returns false as well as explanation why date range is invalid.
func areDatesValid(startDate, endDate time.Time) (bool, string) {
	for i, r := range dateValidationRules {
		if !r(startDate, endDate) {
			return false, fmt.Sprintf("validation error: %s", i)
		}
	}

	return true, ""
}

func (n *NasaPicturesHandler) obtainStringDatesFromRequest(r *http.Request) (string,string) {
	sDate := r.URL.Query().Get(startDateParam)
	eDate := r.URL.Query().Get(endDateParam)

	if len(sDate) == 0 || len(eDate) == 0 {
		n.err = customError.Server{Message: "start_date or end_date need to passed as query parameters"}
		return "", ""
	}

	return sDate, eDate
}

func (n *NasaPicturesHandler) parseToDates(sDate, eDate string) (time.Time, time.Time) {
	if n.err != nil {
		return time.Time{}, time.Time{}
	}

	start, err := time.Parse("2006-01-02", sDate)
	end, err := time.Parse("2006-01-02", eDate)
	if err != nil {
		n.err = err
	}

	return start, end
}
