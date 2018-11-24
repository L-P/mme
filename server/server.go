package server

import (
	"log"
	"net/http"
	"time"

	"github.com/L-P/mme/rom"
)

// Server serves ROM info over HTTP.
type Server struct {
	httpServer *http.Server
	rom        *rom.View
}

// New creates a new Server
func New(view *rom.View) *Server {
	s := &Server{
		rom: view,
		httpServer: &http.Server{
			Addr:         "127.0.0.1:3001",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}

	s.setupRoutes()

	return s
}

// ListenAndServe serves the app over HTTP
func (s *Server) ListenAndServe() error {
	log.Printf("Starting server at %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) setupRoutes() {
	http.HandleFunc("/", s.indexHandler)
}

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/"))
}
