package main

import (
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {

	rdb := CreateDBConnection(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	sockethandler := NewSocketHandler(rdb)
	roomhandler := NewRoomHandler(rdb)

	mux := http.NewServeMux()
	mux.Handle("/ws/", sockethandler)
	mux.Handle("/room/", roomhandler)

	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}
