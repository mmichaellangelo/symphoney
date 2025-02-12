package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type RoomHandler struct {
	db *DB
}

var (
	RoomRE       = regexp.MustCompile(`^\/room\/?$`)
	RoomREWithID = regexp.MustCompile(`^\/room\/(\d+)\/?$`)
)

func NewRoomHandler(db *DB) *RoomHandler {
	return &RoomHandler{db: db}
}

func (h *RoomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	url := r.URL.Path
	switch {
	// Create Room
	case r.Method == http.MethodPost && RoomRE.MatchString(url):
		roomID, err := h.db.CreateRoom()
		if err != nil {
			http.Error(w, fmt.Sprintf("error creating room: %v", err), http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(map[string]string{"roomID": roomID})
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(data)
		if err != nil {
			fmt.Println("error sending data: %w", err)
		}
		break

	// Get All Rooms
	case r.Method == http.MethodGet && RoomRE.MatchString(url):

	}
}

func (h *RoomHandler) CreateRoom() {

}
