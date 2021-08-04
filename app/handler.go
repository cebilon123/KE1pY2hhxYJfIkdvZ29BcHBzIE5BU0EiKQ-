package app

import (
	"encoding/json"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/customError"
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
	defer ctx.Done()

	//TODO validate and unit test for example this one and all validations
	sDate := r.URL.Query()[startDateParam]
	eDate := r.URL.Query()[endDateParam]
	if len(sDate) == 0 || len(eDate) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr("Wrong date format").ToByteSlice())
		return
	}


	start,err := time.Parse("2006-01-02",sDate[0])
	end,err := time.Parse("2006-01-02",eDate[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr(err.Error()).ToByteSlice())
		return
	}

	res, err := ip.GetImagesUrls(start,end)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(customError.NewApiErr(err.Error()).ToByteSlice())
		return
	}

	w.WriteHeader(http.StatusOK)
	mRes, _ := json.Marshal(res)
	w.Write(mRes)
}
