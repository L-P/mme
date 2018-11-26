package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/L-P/mme/colormap"
	"github.com/L-P/mme/rom"
	"github.com/gobuffalo/packr"
	"github.com/husobee/vestigo"
)

// Server serves ROM info over HTTP.
type Server struct {
	httpServer *http.Server
	rom        *rom.View
	static     packr.Box
	router     *vestigo.Router
}

// New creates a new Server
func New(view *rom.View) *Server {
	router := vestigo.NewRouter()

	s := &Server{
		rom: view,
		httpServer: &http.Server{
			Addr:         "127.0.0.1:8064",
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
			Handler:      router,
		},
		static: packr.NewBox("../front/dist"),
		router: router,
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
	s.router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin: []string{
			"http://localhost:3000",
			"http://localhost:8064",
			"http://localhost:8080",
		},
	})

	s.router.Get("/api/rom", s.romHandler)
	s.router.Get("/api/colormap", s.colormapHandler())
	s.router.Get("/api/messages", s.messagesHandler)

	s.router.Get("/api/rooms/:start", s.roomDetailHandler)
	s.router.Get("/api/scenes/:start", s.sceneDetailHandler)
	s.router.Get("/api/scenes", s.scenesHandler)

	s.router.Get("/api/files/:start", s.fileDataHandler)
	s.router.Get("/api/files", s.filesHandler)

	// Static and generated files
	s.router.Get("/", s.indexHandler)
	s.router.Get("/:route", s.indexHandler)
	s.router.Get("/favicon.ico", s.faviconHandler)
	s.router.Handle("/assets/:file", s.handleStatic())
	s.router.Handle("/_/:file", s.handleStatic())
	s.router.Handle("/_/:type/:file", s.handleStatic())
}

func (s *Server) handleStatic() http.Handler {
	return s.addCacheHeaders(http.FileServer(s.static))
}

// Catch-all to index to allow for Vue URIs
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := s.static.Find("index.html")
	w.Write(b)
}

// Required to be at the root, so here it is.
func (s *Server) faviconHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := s.static.Find("favicon.ico")
	w.Write(b)
}

func (s *Server) colormapHandler() func(w http.ResponseWriter, r *http.Request) {
	var cmap bytes.Buffer

	return func(w http.ResponseWriter, r *http.Request) {
		if cmap.Len() <= 0 {
			err := colormap.Generate(&cmap, s.rom)
			if err != nil {
				log.Fatal(err)
			}
		}

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

func (s *Server) messagesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(s.rom.Messages)
}

func (s *Server) romHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	rom := s.rom.GetROM()
	team, date := rom.ParseBuild()

	enc.Encode(map[string]interface{}{
		"Name":       string(rom.Name[:]),
		"CRC1":       fmt.Sprintf("%08X", rom.CRC1),
		"CRC2":       fmt.Sprintf("%08X", rom.CRC2),
		"Build team": team,
		"Build date": date,
	})
}
