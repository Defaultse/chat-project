package repository

import (
	"chat-project-go/internal/datastruct"
	"database/sql"
	"fmt"
)

type ChatRepositoryContract interface {
	GetAllChatsByUserId(userId string) ([]datastruct.ChatsLastMsgs, error)
	CreateConversation(firstUserID string, secondUserID string) (int64, error)
	CreateMessage(from_user int64, conversation_id int64, text string) error
}

type ChatRepository struct {
	db func() *sql.DB
}

func NewChatRepository(db func() *sql.DB) ChatRepositoryContract {
	return ChatRepository{db: db}
}

func (u ChatRepository) CreateConversation(firstUserID string, secondUserID string) (int64, error) {
	// !!!!!!!!Здесь либо нужны транзакции, либо триггер чтобы автоматом создавалась табличка
	var conversationId int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`INSERT INTO dbo.conversation DEFAULT VALUES; SELECT SCOPE_IDENTITY();`)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	if err := u.db().QueryRow(query).Scan(&conversationId); err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	query = fmt.Sprintf(`INSERT INTO users_conversations(conversation_id, user_id) VALUES(%d, %s),(%d, %s);`, conversationId, firstUserID, conversationId, secondUserID)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	if err := u.db().QueryRow(query).Scan(&conversationId); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return conversationId, nil
}

func (u ChatRepository) GetAllChatsByUserId(userId string) ([]datastruct.ChatsLastMsgs, error) {
	var msgs []datastruct.ChatsLastMsgs

	query := fmt.Sprintf(`SELECT * FROM (
		SELECT ROW_NUMBER() OVER (PARTITION BY message.conversation_id ORDER BY message.created_at) AS LastMessage, message.conversation_id, message.message_id, message.from_user, message.text, message.created_at
		FROM users_conversations
		JOIN message
		ON users_conversations.conversation_id = message.conversation_id
		WHERE users_conversations.user_id=%s) AS a WHERE a.LastMessage = 1`, userId)

	rows, err := u.db().Query(query)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var msg *datastruct.ChatsLastMsgs
		err = rows.Scan(&msg)

		if err != nil {
			fmt.Println(err)
			continue
		}

		msgs = append(msgs, *msg)
	}

	return msgs, nil
}

func (u ChatRepository) CreateMessage(from_user int64, conversation_id int64, text string) error {
	var err error

	query := fmt.Sprintf(`INSERT INTO dbo.message(from_user, text, conversation_id) VALUES(%d, %d, %s)`, from_user, conversation_id, text)

	if _, err := u.db().Exec(query); err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
