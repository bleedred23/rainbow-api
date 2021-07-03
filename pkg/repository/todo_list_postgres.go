package repository

import (
	rainbow_api "github.com/bleedred23/rainbow-api"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r* TodoListPostgres) Create(userId int, list rainbow_api.TodoList) (int, error) {

}
