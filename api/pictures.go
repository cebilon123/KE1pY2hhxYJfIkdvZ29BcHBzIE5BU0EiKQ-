package api

import (
	"net/http"
	"time"

	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/config"
	"github.com/gin-gonic/gin"
)

type PictureQuery struct {
	From time.Time `form:"from" time_format:"2006-01-02" time_utc:"1"`
	To   time.Time `form:"to" time_format:"2006-01-02" time_utc:"1"`
}

func GetPictures(ctx *gin.Context, cfg config.ConfigReader) {
	var query PictureQuery
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, error("Invalid query params"))
		return
	}
	println("going")
}
