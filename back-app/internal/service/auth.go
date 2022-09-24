package service

import (
	"chat-project-go/internal/dto"
	"chat-project-go/internal/repository"
	"crypto/md5"
	"fmt"
	"strconv"
)

type AuthService interface {
	SignUp(user *dto.User) (*int64, error)
	SignIn(email, password string) (*string, error)
	Logout(userID int64) error
}

type authService struct {
	userRepository repository.UserRepositoryContract
	tokenManager   TokenManager
}

func NewAuthService(tokenManager TokenManager, userRepository repository.UserRepositoryContract) AuthService {
	return &authService{
		userRepository: userRepository,
		tokenManager:   tokenManager,
	}
}

func (a *authService) SignUp(user *dto.User) (*int64, error) {
	user.PasswordHash = fmt.Sprintf("%x", md5.Sum([]byte(user.PasswordHash)))

	id, err := a.userRepository.CreateUser(*user)

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (a *authService) SignIn(email, reqPassword string) (*string, error) {
	reqPassword = fmt.Sprintf("%x", md5.Sum([]byte(reqPassword)))

	password, err := a.userRepository.GetUserPasswordByEmail(email)
	if err != nil {
		return nil, err
	}

	if password != reqPassword {
		return nil, fmt.Errorf("passwords don't match")
	} else {
		userID, err := a.userRepository.GetUserIdByEmail(email)
		if err != nil {
			return nil, err
		}

		jwt, err := a.tokenManager.NewJWT(strconv.Itoa(int(userID)))
		if err != nil {
			return nil, err
		}

		return &jwt, nil
	}
}

func (a *authService) Logout(userID int64) error {
	_, err := a.tokenManager.NewJWT(strconv.Itoa(int(userID)))
	if err != nil {
		return err
	}
	return nil
}
