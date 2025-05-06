package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"shortlinks/internal/user"
)

type AuthService struct {
	UseRepository *user.UserRepository
}

func NewAuthService(repository *user.UserRepository) *AuthService {
	return &AuthService{repository}
}

func (service *AuthService) Login(email, password string) (string, error) {
	existedUser, _ := service.UseRepository.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New(ErrorWrongCredentials)
	}
	err := bcrypt.CompareHashAndPassword([]byte(existedUser.Password), []byte(password))

	if err != nil {
		return "", errors.New(ErrorWrongCredentials)
	}
	return existedUser.Email, nil
}

func (service *AuthService) Register(email, password, name string) (string, error) {
	existedUser, _ := service.UseRepository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New(ErrUserExists)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}
	_, err = service.UseRepository.Create(user)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
