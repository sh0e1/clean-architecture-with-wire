package database

import (
	"context"

	"github.com/sh0e1/wire/domain"
)

func NewTodoRepository(sqlHander SQLHandler) *TodoRepository {
	return &TodoRepository{
		SQLHandler: sqlHander,
	}
}

type TodoRepository struct {
	SQLHandler
}

func (r *TodoRepository) Store(ctx context.Context, todo *domain.Todo) error {
	const sqlstr = `insert into todos (id, title) values(?, ?)`
	_, err := r.ExecuteContext(ctx, sqlstr, todo.ID, todo.Title)
	return err
}

func (r *TodoRepository) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	todo := &domain.Todo{}
	const sqlstr = `select * from todos where id = ?`
	if err := r.QueryRowContext(ctx, sqlstr, id).Scan(&todo.ID, &todo.Title); err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *TodoRepository) List(ctx context.Context) ([]*domain.Todo, error) {
	const sqlstr = `select * from todos`
	rows, err := r.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, err
	}
	// nolint:errcheck
	defer rows.Close()

	var todos = make([]*domain.Todo, 0)
	for rows.Next() {
		todo := &domain.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()
}
