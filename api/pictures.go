package api

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/collector"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/config"
	"github.com/gin-gonic/gin"
)

type pictureQuery struct {
	From time.Time `form:"from" time_format:"2006-01-02" time_utc:"1"`
	To   time.Time `form:"to" time_format:"2006-01-02" time_utc:"1"`
}

func (pq *pictureQuery) validate() url.Values {
	errs := url.Values{}

	if pq.To.Before(pq.From) {
		errs.Add("from", "from need to be before to")
	}

	if pq.To.After(time.Now().AddDate(0, 0, -1)) {
		errs.Add("to", "to can't be after yesterday")
	}

	if pq.From.Before(time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)) {
		errs.Add("from", "from need to be at least 2015-01-01")
	}

	return errs
}

func GetPictures(ctx *gin.Context, cfg config.ConfigReader, imgCollector collector.ImageCollector) {
	var query = &pictureQuery{}
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, error("Invalid query params"))
		return
	}

	if errs := query.validate(); len(errs) > 0 {
		strErr := ""
		for param, errors := range errs {
			strErr += param + ": " + strings.Join(errors, "") + ` `
		}
		ctx.JSON(http.StatusUnprocessableEntity, error(strErr))
		return
	}

	urls, err := imgCollector.GetImages(query.From, query.To)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, error(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}
