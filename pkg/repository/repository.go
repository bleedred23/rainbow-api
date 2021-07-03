package repository

import (
	"github.com/bleedred23/rainbow-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rainbow_api.User) (int, error)
	GetUser(username, password string) (rainbow_api.User, error)
}

type TodoList interface {
	Create(userId int, list rainbow_api.TodoList) (int, error)
}

type TodoItem interface {

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),

	}
}
