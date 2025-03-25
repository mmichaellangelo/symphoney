package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/redis/go-redis/v9"
)

type DB struct {
	db *redis.Client
}

const (
	letters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	uuid_length = 4
)

const (
	ROOM_SET = "ROOMS"
)

var ctx = context.Background()

/*
* DATABASE FUNCTIONS
 */
func CreateDBConnection(options *redis.Options) *DB {
	rdb := redis.NewClient(options)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &DB{db: rdb}
}

/*
* ROOM FUNCTIONS
 */

func (d *DB) CreateRoom() (roomID string, err error) {
	id, err := d.generateRoomID()
	if err != nil {
		return "", fmt.Errorf("error creating room: %w", err)
	}
	return id, nil
}

func (d *DB) generateRoomID() (string, error) {
	code := make([]byte, uuid_length)
	for {
		for i := range code {
			code[i] = letters[rand.Intn(len(letters))]
		}

		codeString := string(code)

		exists, err := d.RoomExists(codeString)
		if err != nil {
			return "", err
		}
		if exists {
			return codeString, nil
		}
	}
}

func (d *DB) RoomExists(id string) (bool, error) {
	res, err := d.db.SIsMember(context.Background(), ROOM_SET, id).Result()
	if err != nil {
		return false, fmt.Errorf("error checking room ID: %w", err)
	}
	return res, nil
}

func (d *DB) GetAllRoomIDs() ([]string, error) {
	rooms, err := d.db.SMembers(context.Background(), ROOM_SET).Result()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (d *DB) AddMember(roomID string) (memberID string, err error) {
	roomExists, err := d.RoomExists(roomID)
	if err != nil {
		return "", fmt.Errorf("error checking if room exists: %w", err)
	}
	if !roomExists {
		return "", fmt.Errorf("error: room does not exist")
	}
	memberID = ""
	i := 1
	for {
		memberID = fmt.Sprintf("member%d", i)

		exists, err := d.db.SIsMember(ctx, roomID, memberID).Result()
		if err != nil {
			return "", fmt.Errorf("error checking if is member: %w", err)
		}
		if !exists {
			break
		}
		i++
	}

	err = d.db.SAdd(ctx, roomID, memberID).Err()
	if err != nil {
		return "", fmt.Errorf("error adding member to room: %w", err)
	}
	return memberID, nil
}

func (d *DB) GetRoomCardinality(roomID string) (int64, error) {
	exists, err := d.RoomExists(roomID)
	if err != nil {
		return -1, err
	}
	if !exists {
		return -1, fmt.Errorf("error: room does not exist")
	}

	card, err := d.db.SCard(ctx, roomID).Result()
	if err != nil {
		return -1, fmt.Errorf("error getting room cardinality: %w", err)
	}
	return card, nil
}

func (d *DB) GetRoomMemberIDs(roomID string) ([]string, error) {
	exists, err := d.RoomExists(roomID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("error: room does not exist")
	}

	members, err := d.db.SMembers(ctx, roomID).Result()
	if err != nil {
		return nil, fmt.Errorf("error getting room cardinality: %w", err)
	}
	return members, nil
}

// func (d *DB) GetRoomMemberDataAll(roomID string) (map[string]interface{}, error) {
// 	exists, err := d.RoomExists(roomID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if !exists {
// 		return nil, fmt.Errorf("error: room does not exist")
// 	}

// }

func (d *DB) RemoveMember(roomID string, memberID string) error {
	fmt.Printf("Removing member %s from room %s\n", memberID, roomID)
	// memberHashkey := GetMemberHashKey(memberID, roomID)
	// err := d.db.HDel(ctx, memberHashkey).Err()
	// if err != nil {
	// 	return fmt.Errorf("error deleting member hash: %w", err)
	// }
	err := d.db.SRem(ctx, roomID, memberID).Err()
	if err != nil {
		return fmt.Errorf("error removing member from set: %w", err)
	}
	return nil
}

func (d *DB) PublishData(roomID string, memberID string, data json.RawMessage) error {
	// hashKey := GetMemberHashKey(memberID, roomID)
	// err := d.db.HSet(ctx, hashKey, "data", data).Err()
	msg := Message{
		RoomID:   &roomID,
		MemberID: &memberID,
		Data:     &data,
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
