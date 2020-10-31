package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sh0e1/wire/domain"
)

func NewTodoInteractor(todoRopo TodoRepository) *TodoInteractor {
	return &TodoInteractor{
		TodoRepository: todoRopo,
	}
}

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (i *TodoInteractor) Create(ctx context.Context, todo *domain.Todo) error {
	todo.ID = uuid.New().String()
	return i.TodoRepository.Store(ctx, todo)
}

func (i *TodoInteractor) GetByID(ctx context.Context, id string) (*domain.Todo, error) {
	return i.TodoRepository.GetByID(ctx, id)
}

func (i *TodoInteractor) List(ctx context.Context) ([]*domain.Todo, error) {
	return i.TodoRepository.List(ctx)
}

type TodoRepository interface {
	Store(ctx context.Context, todo *domain.Todo) error
	GetByID(ctx context.Context, id string) (*domain.Todo, error)
	List(ctx context.Context) ([]*domain.Todo, error)
}
