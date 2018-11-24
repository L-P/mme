package server

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/L-P/mme/colormap"
	"github.com/L-P/mme/rom"
	"github.com/gobuffalo/packr"
)

// Server serves ROM info over HTTP.
type Server struct {
	httpServer *http.Server
	rom        *rom.View
	static     packr.Box
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
		static: packr.NewBox("../static"),
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
	http.Handle("/", http.FileServer(s.static))
	http.HandleFunc("/api/colormap", s.colormapHandler())
}

func (s *Server) colormapHandler() func(w http.ResponseWriter, r *http.Request) {
	var cmap bytes.Buffer
	err := colormap.Generate(&cmap, s.rom)
	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(cmap.Bytes())
	}
}
