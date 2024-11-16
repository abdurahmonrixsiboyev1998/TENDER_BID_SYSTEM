package service

import (
	"context"
	"errors"
	"tender_bid_system/model"
	"tender_bid_system/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, user *model.User) (model.User, error) {
	existingUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err == nil && existingUser.ID > 0 {
		return model.User{}, errors.New("email already exists")
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) LoginUser(ctx context.Context, email, password string) (string, error) {
	token, err := s.repo.Login(ctx, email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
