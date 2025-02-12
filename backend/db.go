package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	db *redis.Client
}

var ctx = context.Background()

func CreateDBConnection(options *redis.Options) *DB {
	rdb := redis.NewClient(options)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return &DB{db: rdb}
}

func (d *DB) CreateRoom() (roomID string, err error) {
	newRoomNumber, err := d.db.Incr(ctx, "room_counter").Result()
	if err != nil {
		return "", fmt.Errorf("error incrementing room counter: %w", err)
	}
	roomID = fmt.Sprintf("room%d", newRoomNumber)
	return roomID, nil
}

func (d *DB) GetAllRoomIDs() ([]string, error) {
	var cursor uint64
	var rooms []string
	for {
		var keys []string
		var err error
		keys, cursor, err = d.db.Scan(ctx, cursor, "room*", 10).Result()
		if err != nil {
			return nil, fmt.Errorf("error scanning db: %w", err)
		}
		rooms = append(rooms, keys...)
		if cursor == 0 {
			break
		}
	}
	return rooms, nil
}

func (d *DB) AddMember(roomID string) (memberID string, err error) {
	memberNumber, err := d.db.SCard(ctx, roomID).Result()
	if err != nil {
		return "", fmt.Errorf("error getting cardinality of room: %w", err)
	}
	memberID = fmt.Sprintf("member%d", memberNumber+1)
	_, err = d.db.SAdd(ctx, roomID, memberID).Result()
	if err != nil {
		return "", fmt.Errorf("error adding member to room: %w", err)
	}
	memberHashKey := fmt.Sprintf("%s:%s", roomID, memberID)
	_, err = d.db.HSet(ctx, memberHashKey, "x", 0, "y", 0).Result()
	return memberID, nil
}

type Message struct {
	RoomID   string `json:"roomID"`
	MemberID string `json:"memberID"`
	Data     string `json:"data"`
}

func (d *DB) PublishData(roomID string, memberID string, data string) error {
	hashKey := GetMemberHashKey(memberID, roomID)
	err := d.db.HSet(ctx, hashKey, "data", data).Err()
	msg := Message{
		RoomID:   roomID,
		MemberID: memberID,
		Data:     data,
	}
	msgJson, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}
	d.db.Publish(ctx, roomID, msgJson)
	return nil
}

func (d *DB) SubscribeToChannel(roomID string) *redis.PubSub {
	pubsub := d.db.Subscribe(ctx, roomID)
	return pubsub
}

func GetMemberHashKey(memberID string, roomID string) string {
	return fmt.Sprintf("%s:%s", roomID, memberID)
}
