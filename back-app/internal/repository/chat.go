package repository

import (
	"chat-project-go/internal/datastruct"
	"database/sql"
	"fmt"
)

type ChatRepositoryContract interface {
	GetAllChatsByUserId(userId int) ([]datastruct.AllChats, error)
}

type ChatRepository struct {
	db func() *sql.DB
}

func NewChatRepository(db func() *sql.DB) ChatRepositoryContract {
	return ChatRepository{db: db}
}

// SELECT DISTINCT(chat_id) FROM chat_and_users WHERE chat_id IN (SELECT chat_id FROM chat_and_users WHERE user_id=2014);

// SELECT DISTINCT (chat_id), id, user_id, line_text, created_at FROM chat_line
//                                                               WHERE chat_id in
//                                                                     (SELECT DISTINCT(chat_id) FROM chat_and_users WHERE chat_id IN
//                                                                                                                         (SELECT chat_id FROM chat_and_users WHERE user_id=2014)
//                                                                                                                   )
//                                                               order by created_at desc;

func (u ChatRepository) GetAllChatsByUserId(userId int) ([]datastruct.AllChats, error) {
	var id int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(``, userId)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer stmt.Close()

	if err := u.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}
