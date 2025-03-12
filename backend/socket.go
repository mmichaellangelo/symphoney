package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

type SocketHandler struct {
	db *DB
}

var (
	WSRoomClientRE = regexp.MustCompile(`^\/ws\/room\/([a-zA-Z0-9_]+)\/client\/?$`)
	WSRoomServerRE = regexp.MustCompile(`^\/ws\/room\/([a-zA-Z0-9_]+)\/server\/?$`)
)

type BaseMsg struct {
	Type string `json:"type"`
}

type MemberAddedMsg struct {
	Type     BaseMsg
	MemberID string
	RoomID   string
}

type MemberUpdatedMsg struct {
	Type     BaseMsg
	MemberID string
	RoomID   string
	Data     map[string]interface{}
}

type Member struct {
	ID string `json:"id"`
}

type Room struct {
	ID      string
	Members []Member
}

func NewSocketHandler(db *DB) *SocketHandler {
	return &SocketHandler{db: db}
}

func (h *SocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	url := r.URL.Path
	switch {
	// Join Room as Client
	case WSRoomClientRE.MatchString(url):
		h.HandleClient(w, r)
		break
	// Join Room as Server
	case WSRoomServerRE.MatchString(url):
		h.HandleServer(w, r)
		break
	}

}

func (h *SocketHandler) HandleClient(w http.ResponseWriter, r *http.Request) {
	roomID := WSRoomClientRE.FindStringSubmatch(r.URL.Path)[1]
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"symphoney.xyz"},
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer c.CloseNow()
	log.Println("Client Websocket Connection Initialized With " + r.Host)

	memberID, err := h.db.AddMember(roomID)
	if err != nil {
		fmt.Println("error adding member to room: %w", err)
		c.CloseNow()
	}
	defer h.db.RemoveMember(roomID, memberID)

	for {
		_, rdr, err := c.Reader(context.Background())
		if err != nil {
			log.Println("Error reading from WebSocket:", err)
			break
		}

		msg := make([]byte, 1024)
		n, err := rdr.Read(msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		data := string(msg[:n])
		log.Printf("%s\n", data)
		h.db.PublishData(roomID, memberID, data)
	}
}

func (h *SocketHandler) HandleServer(w http.ResponseWriter, r *http.Request) {
	roomID := WSRoomServerRE.FindStringSubmatch(r.URL.Path)[1]
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"symphoney.xyz"},
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer c.CloseNow()
	log.Println("Server Websocket Connection Initialized With " + r.Host)
	pubsub := h.db.db.Subscribe(ctx, roomID)
	defer pubsub.Unsubscribe(ctx, roomID)
	defer pubsub.Close()
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
		wsjson.Write(ctx, c, msg.Payload)
	}

}

func testingCodeDoNotRun(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		OriginPatterns: []string{"symphoney.xyz"},
	})
	if err != nil {
		log.Println(err)
		return
	}
	defer c.CloseNow()
	log.Println("Websocket initialized with " + r.Host)

	for {
		_, rdr, err := c.Reader(context.Background())
		if err != nil {
			log.Println("Error reading from WebSocket:", err)
			break
		}

		msg := make([]byte, 1024)
		n, err := rdr.Read(msg)
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("%s\n", string(msg[:n]))
	}
}
