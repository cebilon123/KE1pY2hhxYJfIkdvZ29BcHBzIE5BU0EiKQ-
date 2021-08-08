package server

import (
	"fmt"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/customError"
	"net/http"
	"os"
	"strconv"
)

// Server represents server instance, which stores all fields
// which are needed to start application correctly
type Server struct {
	port   string
	config config
}

// config contains all configuration
type config struct {
	nasaApiKey            string
	maxConcurrentApiCalls int8 // there is no need for bigger value
}

const (
	defaultPort = ":8080"

	ApiKey                = "Api_Key"
	MaxConcurrentApiCalls = "Max_Concurrent_Api_Calls"
)

// NewServer creates new server instance,
// with port being set to default value as 8080.
func NewServer() *Server {
	return &Server{port: defaultPort}
}

// WithPort is server builder function which sets server port
// as desired value. Default port: 8080.
func (s *Server) WithPort(port string) *Server {
	s.port = port

	return s
}

// WithConfig creates server config.
func (s *Server) WithConfig(nasaApiKey string, maxConcurrentApiCalls int8) *Server {
	s.config.nasaApiKey = nasaApiKey
	s.config.maxConcurrentApiCalls = maxConcurrentApiCalls

	return s
}

// AddHandler adds handler to the server.
func (s *Server) AddHandler(pattern string, handler func(w http.ResponseWriter, r *http.Request)) *Server {
	fmt.Printf("Endpoint: %s\n", pattern)
	http.HandleFunc(pattern, handler)

	return s
}

// Start starts http server. It validates server configuration as well.
// If configuration is invalid or there is problem with starting server,
// it returns error.
func (s *Server) Start() error {
	isValid, field := isValidServer(s)
	if !isValid {
		return customError.Server{
			Message: fmt.Sprintf("configuration isn't valid, field: %s", field),
		}
	}

	setEnvVars(s)

	fmt.Printf("Server started on port %s\n", s.port)
	if err := http.ListenAndServe(s.port, nil); err != nil {
		return err
	}

	return nil
}

// isValidServer checks if server configuration is valid.
// Returns bool whenever server is valid, and string which is
// the name of first found invalid field whenever function returns false.
func isValidServer(s *Server) (bool, string) {
	if len(s.config.nasaApiKey) == 0 {
		return false, "nasaApiKey"
	}
	if s.config.maxConcurrentApiCalls <= 0 {
		return false, "maxConcurrentApiCalls"
	}

	return true, ""
}

// setEnvVars sets env variables for application.
func setEnvVars(s *Server) {
	os.Setenv(ApiKey, s.config.nasaApiKey)
	os.Setenv(MaxConcurrentApiCalls, strconv.Itoa(int(s.config.maxConcurrentApiCalls)))
}
