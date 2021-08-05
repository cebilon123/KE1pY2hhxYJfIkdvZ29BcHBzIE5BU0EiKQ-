package config

import "flag"

// Startup represents configuration which should be build on startup
// with use of flags
type Startup struct {
	ApiKey                    string
	Port                      string
	MaxConcurrentNasaRequests int
}

const (
	apiKeyName = "API_KEY"
	defaultApiKey = "DEMO_KEY"
	portName = "PORT"
	defaultPort = "8080"
	concurrentRequests = "CONCURRENT_REQUESTS"
	defaultConcurrentRequests = 5
)

func NewStartup() *Startup {
	var apiKey, port string
	var maxConcurrentNasaRequests int

	flag.StringVar(&apiKey, apiKeyName, defaultApiKey, "Specify api key used to communicate with NASA api")
	flag.StringVar(&port, portName, defaultPort, "Specify port on which application should listen and serve")
	flag.IntVar(&maxConcurrentNasaRequests, concurrentRequests, defaultConcurrentRequests, "Specify max number of concurrent requests to NASA api")

	return &Startup{ApiKey: apiKey, Port: port, MaxConcurrentNasaRequests: maxConcurrentNasaRequests}
}


