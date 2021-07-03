package service

import (
	rainbow_api "github.com/bleedred23/rainbow-api"
	"github.com/bleedred23/rainbow-api/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s* TodoListService) Create(userId int, list rainbow_api.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
