package main

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

type Message struct {
	// Type     MessageType     `json:"type"`
	RoomID   *string `json:"roomID,omitempty"`
	MemberID *string `json:"memberID,omitempty"`
	Data     *string `json:"data"`
}
