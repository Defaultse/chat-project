package datastruct

import "time"

type ChatsLastMsgs struct {
	ConversationID int64     `json:"conversation_id"`
	MessageID      int64     `json:"message_id"`
	FromUser       int64     `json:"from_user"`
	Text           string    `json:"text"`
	CreatedAt      time.Time `json:"created_at"`
}
