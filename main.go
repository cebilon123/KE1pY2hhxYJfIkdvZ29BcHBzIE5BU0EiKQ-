package main

import "github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/server"

func main() {
	if err := server.NewServer().
		WithPort(":9999").
		WithConfig("tN5MEJyrF1HZKVGrUvPrPiIM44vcm0ByOp0UqWMW",5).
		AddHandler("/pictures", server.NasaPicturesHandler).
		Start()
	err != nil {
		panic(err)
	}
}
