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
			Addr:         "127.0.0.1:8064",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		static: packr.NewBox("../front/dist"),
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
	// Static and generated files
	http.HandleFunc("/", s.indexHandler)
	http.HandleFunc("/favicon.ico", s.faviconHandler)
	http.Handle("/assets/", s.addCacheHeaders(http.FileServer(s.static)))
	http.Handle("/_/", s.addCacheHeaders(http.FileServer(s.static)))

	http.HandleFunc("/api/colormap", s.colormapHandler())
}

// Catch-all to index to allow for Vue URIs
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	b, err := s.static.Find("index.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

// Required to be at the root, so here it is.
func (s *Server) faviconHandler(w http.ResponseWriter, r *http.Request) {
	b, err := s.static.Find("favicon.ico")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
}

func (s *Server) colormapHandler() func(w http.ResponseWriter, r *http.Request) {
	var cmap bytes.Buffer
	go func() {
		err := colormap.Generate(&cmap, s.rom)
		if err != nil {
			log.Fatal(err)
		}
	}()

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.Write(cmap.Bytes())
	}
}

func (s *Server) addCacheHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=31536000, public, immutable")
		h.ServeHTTP(w, r)
	}
}
