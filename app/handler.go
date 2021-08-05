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

// dateValidationRules are rules used to validate dates in AreDatesValid() function
// for sake of simplicity here can be passed arbitrary number of time arguments
// bue first one should be startDate (which is earlier than endDate)
var dateValidationRules = map[string]func(dates ...time.Time) bool{
	"start_date should be earlier than end_date" : func(dates ...time.Time) bool {
		return dates[0].Before(dates[1])
	},
}

// NasaPicturesHandler handles image url request
func NasaPicturesHandler(w http.ResponseWriter, r *http.Request, ip domain.ImageProvider) {
	ctx := r.Context()
	defer ctx.Done()

	sDate := r.URL.Query()[startDateParam]
	eDate := r.URL.Query()[endDateParam]

	//TODO change func args and parse
	datesValid, msg := AreDatesValid(sDate, eDate)
	if  {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr()
		return
	}

	start, err := time.Parse("2006-01-02", sDate[0])
	end, err := time.Parse("2006-01-02", eDate[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr(err.Error()).ToByteSlice())
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

// AreDatesValid Checks if date range is valid. Returns true whenever date is valid
// and returns false as well as explanation why date range is invalid.
func AreDatesValid(startDate, endDate []time.Time) (bool, string) {
	if len(startDate) == 0 || len(startDate) == 0 {
		return false, "start_date or end_date need to passed as query parameters"
	}

	for i, r := range dateValidationRules {
		if !r(startDate[0], endDate[0]) {
			return false, fmt.Sprintf("validation error: %s",i)
		}
	}

	return true, ""
}
