package server

import (
	"fmt"
	"github.com/cebilon123/KE1pY2hhxYJfIkdvZ29BcHBzIE5BU0EiKQ-/customError"
	"net/http"
)

type Server struct {
	port   string
	config config
}

type config struct {
	nasaApiKey string
}

const (
	defaultPort = ":8080"
)

func NewServer() *Server {
	return &Server{port: defaultPort}
}

func (s *Server) WithPort(port string) *Server {
	s.port = port

	return s
}

func (s *Server) Start() error {
	isValid, field := isValidServer(s)
	if !isValid {
		return customError.Server{
			Message: fmt.Sprintf("configuration isn't valid, field: %s", field),
		}
	}

	if err := http.ListenAndServe(s.port, nil); err != nil {
		return err
	}

	return nil
}

// isValidServer checks if server configuration is valid.
// Returns bool whenever server is valid, and string which is
// the name of first found invalid field whenever function returns false
func isValidServer(s *Server) (bool, string) {
	if len(s.config.nasaApiKey) == 0 {
		return false, "nasaApiKey"
	}

	return true, ""
}
