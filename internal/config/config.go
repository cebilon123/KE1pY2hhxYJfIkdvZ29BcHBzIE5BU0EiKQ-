package config

type ConfigReader interface {
	Port() string
	ApiKey() string
	ConcurrentRequests() int
}

type config struct {
	port               string
	apiKey             string
	concurrentRequests int
}

func NewConfig(port string, apiKey string, concurrentRequests int) ConfigReader {
	return config{
		port:               port,
		apiKey:             apiKey,
		concurrentRequests: concurrentRequests,
	}
}

func (c config) Port() string {
	return c.port
}

func (c config) ApiKey() string {
	return c.apiKey
}

func (c config) ConcurrentRequests() int {
	return c.concurrentRequests
}
