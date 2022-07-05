package UberDi

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	config        *Config
	personService *PersonService
}

func NewServer(config *Config, service *PersonService) *Server {
	return &Server{
		config:        config,
		personService: service,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", s.findPeople)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) findPeople(writer http.ResponseWriter, request *http.Request) {
	people := s.personService.FindAll()
	bytes, _ := json.Marshal(people)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(bytes)
}
