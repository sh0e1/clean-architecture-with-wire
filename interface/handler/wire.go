//+build wireinject

package handler

import (
	"github.com/google/wire"

	"github.com/sh0e1/wire/interface/database"
	"github.com/sh0e1/wire/usecase"
)

func InitializeTodoHandler(sqlHander database.SQLHandler) *TodoHandler {
	wire.Build(
		NewTodoHandler,
		usecase.NewTodoInteractor,
		wire.Bind(new(usecase.TodoRepository), new(*database.TodoRepository)),
		database.NewTodoRepository,
	)
	return &TodoHandler{}
}
