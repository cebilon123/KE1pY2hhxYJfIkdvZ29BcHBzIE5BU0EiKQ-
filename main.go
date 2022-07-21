package main

import (
	"fmt"

	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/config"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/internal/environment"
)

var configuration config.ConfigReader

func main() {
	env := environment.LoadEnv()
	configuration = config.NewConfig(env.Port(), env.ApiKey(), env.ConcurrentRequests())
	fmt.Printf("%v", configuration)
}
