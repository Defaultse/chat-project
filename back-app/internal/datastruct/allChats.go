package datastruct

import "time"

type AllChats struct {
	ChatId    int64     `json:"chat_id"`
	Id        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	LineText  string    `json:"line_text"`
	CreatedAt time.Time `json:"created_at"`
}
