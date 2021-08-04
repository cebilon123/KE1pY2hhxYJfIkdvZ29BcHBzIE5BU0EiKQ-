package customError

import "fmt"

type Server struct {
	Message string
}

func (s Server) Error() string {
	return fmt.Sprintf("Server: %s", s.Message)
}
