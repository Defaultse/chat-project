package repository

import (
	"chat-project-go/internal/dto"
	"database/sql"
	"fmt"
)

type UserRepositoryContract interface {
	CreateUser(user dto.User) (int64, error)
	GetUser(id int64) (*dto.User, error)
	GetUserPasswordByEmail(email string) (string, error)
	GetUserIdByEmail(email string) (int64, error)
}

type UserRepository struct {
	db func() *sql.DB
}

func NewUserRepository(db func() *sql.DB) UserRepositoryContract {
	return UserRepository{db: db}
}

func (u UserRepository) CreateUser(user dto.User) (int64, error) {
	var id int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`INSERT INTO dbo.users (username, name, surname, email, phone, password_hash, role) 
		VALUES ('%s', '%s', '%s', '%s', '%v', '%s', '%s'); 
		SELECT SCOPE_IDENTITY()`,
		user.Username, user.Name, user.Surname, user.Email, user.Phone, user.PasswordHash, user.UserType)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	if err := u.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	return id, nil
}

func (u UserRepository) GetUser(id int64) (*dto.User, error) {
	return nil, nil
}

func (u UserRepository) GetUserPasswordByEmail(email string) (string, error) {
	var passwordHash string
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`SELECT password_hash FROM dbo.users WHERE email='%s'`, email)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return "", err
	}

	defer stmt.Close()

	if err := u.db().QueryRow(query).Scan(&passwordHash); err != nil {
		fmt.Println(err)
		return "", err
	}

	return passwordHash, nil
}

func (u UserRepository) GetUserIdByEmail(email string) (int64, error) {
	var id int64
	var stmt *sql.Stmt
	var err error

	query := fmt.Sprintf(`SELECT id FROM dbo.users WHERE email='%s'`, email)

	if stmt, err = u.db().Prepare(query); err != nil {
		fmt.Println(err)
		return 0, err
	}

	defer stmt.Close()

	if err := u.db().QueryRow(query).Scan(&id); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
