package main

import "encoding/json"

type MessageType int

const (
	MessageTypeRoomCreate MessageType = iota
	MessageTypeRoomRead
	MessageTypeRoomUpdate
	MessageTypeRoomDelete
	MessageTypeMemberCreate
	MessageTypeMemberRead
	MessageTypeMemberUpdate
	MessageTypeMemberDelete
)

type Member struct {
	ID     string           `json:"id"`
	Params *json.RawMessage `json:"params"`
}

type Room struct {
	ID     string           `json:"id"`
	Params *json.RawMessage `json:"params"`
}

type Message struct {
	RoomID   *string          `json:"roomID"`
	MemberID *string          `json:"memberID"`
	Data     *json.RawMessage `json:"data"`
}

func (mt MessageType) String() string {
	return [...]string{"RoomCreate", "RoomRead", "RoomUpdate", "RoomDelete", "MemberCreate", "MemberRead", "MemberUpdate", "MemberDelete"}[mt]
}

func (mt MessageType) IsValid() bool {
	switch mt {
	case MessageTypeRoomCreate, MessageTypeRoomRead, MessageTypeRoomUpdate, MessageTypeRoomDelete,
		MessageTypeMemberCreate, MessageTypeMemberRead, MessageTypeMemberUpdate, MessageTypeMemberDelete:
		return true
	}
	return false
}
