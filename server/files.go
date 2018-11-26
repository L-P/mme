package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/husobee/vestigo"
)

func (s *Server) filesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(s.rom.Files)
}

func (s *Server) fileDataHandler(w http.ResponseWriter, r *http.Request) {
	start := vestigo.Param(r, "start")
	i, err := strconv.ParseInt(start, 10, 32)
	if err != nil {
		log.Print(err)
		return
	}

	file, err := s.rom.GetFileByVROMStart(uint32(i))
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Content-Length", fmt.Sprintf("%d", file.Size()))
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment")
	w.Write(file.Data())
}
