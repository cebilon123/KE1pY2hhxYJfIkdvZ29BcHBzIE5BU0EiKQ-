package main

import (
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/app"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/infrastructure"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/server"
	"net/http"
)

func main() {
	if err := server.NewServer().
		WithPort(":9999").
		WithConfig("tN5MEJyrF1HZKVGrUvPrPiIM44vcm0ByOp0UqWMW", 5).
		AddHandler("/pictures", func(w http.ResponseWriter, r *http.Request) {
			app.NasaPicturesHandler(w, r, infrastructure.GetNasaImageProvider())
		}).
		Start(); err != nil {
		panic(err)
	}
}
