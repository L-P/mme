package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/husobee/vestigo"
)

func (s *Server) roomDetailHandler(w http.ResponseWriter, r *http.Request) {
	start, err := strconv.ParseInt(vestigo.Param(r, "start"), 10, 32)
	if err != nil {
		log.Print(err)
		return
	}

	room, err := s.rom.GetRoomByVROMStart(uint32(start))
	if err != nil {
		log.Print(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	enc.Encode(room)
}
