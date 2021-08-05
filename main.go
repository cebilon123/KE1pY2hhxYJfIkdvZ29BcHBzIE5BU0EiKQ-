package main

import (
	"fmt"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/app"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/config"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/infrastructure"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/server"
	"net/http"
	_ "net/http/pprof"
)

//"tN5MEJyrF1HZKVGrUvPrPiIM44vcm0ByOp0UqWMW"
func main() {
	sConfig := config.NewStartup()

	if err := server.NewServer().
		WithPort(fmt.Sprintf(":%s", sConfig.Port)).
		WithConfig(sConfig.ApiKey, int8(sConfig.MaxConcurrentNasaRequests)).
		AddHandler("/pictures", func(w http.ResponseWriter, r *http.Request) {
			app.NasaPicturesHandler(w, r, infrastructure.NewNasaImageProvider())
		}).
		Start(); err != nil {
		panic(err)
	}
}
