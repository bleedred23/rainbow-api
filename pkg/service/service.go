package service

import (
	"github.com/bleedred23/rainbow-api"
	"github.com/bleedred23/rainbow-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user rainbow_api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list rainbow_api.TodoList) (int, error)
}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
