package config

import (
	"os"
	"strconv"
)

// Startup represents configuration which should be build on startup
// with use of flags
type Startup struct {
	ApiKey                    string
	Port                      string
	MaxConcurrentNasaRequests int
}

const (
	apiKeyName                = "API_KEY"
	defaultApiKey             = "DEMO_KEY"
	portName                  = "PORT"
	defaultPort               = "8080"
	concurrentRequests        = "CONCURRENT_REQUESTS"
	defaultConcurrentRequests = 5
)

func NewStartup() *Startup {
	var apiKey, port string
	var maxConcurrentNasaRequests int

	apiKey = os.Getenv(apiKeyName)
	if len(apiKey) == 0 {
		apiKey = defaultApiKey
	}
	port = os.Getenv(portName)
	if len(port) == 0 {
		port = defaultPort
	}
	maxConcurrentNasaRequests, err := strconv.Atoi(os.Getenv(concurrentRequests))
	if err != nil {
		maxConcurrentNasaRequests = defaultConcurrentRequests
	}

	return &Startup{ApiKey: apiKey, Port: port, MaxConcurrentNasaRequests: maxConcurrentNasaRequests}
}
