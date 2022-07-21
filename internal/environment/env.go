package environment

import (
	"os"
	"strconv"
)

// EnvironmentReader is an interface that must be implemented by
// all environment readers. Used for reading values of given
// environment variables.
type EnvironmentReader interface {
	Port() string
	ApiKey() string
	ConcurrentRequests() int
}

type environemnt struct {
	port               string
	apiKey             string
	concurrentRequests int
}

const (
	ApiKeyName             = "API_KEY"
	PortName               = "PORT"
	ConcurrentRequestsName = "CONCURRENT_REQUESTS"

	defaultApiKey             = "DEMO_KEY"
	defaultPort               = "8080"
	defaultConcurrentRequests = "5"
)

// LoadEnv loads environment variables,
// in case of env variables not found
// it assings default values in their place
func LoadEnv() EnvironmentReader {
	apiKey, exist := os.LookupEnv(ApiKeyName)
	if !exist {
		apiKey = defaultApiKey
	}
	port, exist := os.LookupEnv(PortName)
	if !exist {
		port = defaultPort
	}
	port = ":" + port

	concurrentRequests, exist := os.LookupEnv(ConcurrentRequestsName)
	if !exist {
		concurrentRequests = defaultConcurrentRequests
	}

	parsedConcurrentRequests, err := strconv.Atoi(concurrentRequests)
	if err != nil {
		parsedConcurrentRequests = 5
	}

	return environemnt{
		apiKey:             apiKey,
		port:               port,
		concurrentRequests: parsedConcurrentRequests,
	}
}

func (c environemnt) Port() string {
	return c.port
}

func (c environemnt) ApiKey() string {
	return c.apiKey
}

func (c environemnt) ConcurrentRequests() int {
	return c.concurrentRequests
}
