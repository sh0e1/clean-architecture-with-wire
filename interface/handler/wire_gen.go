// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package handler

import (
	"github.com/sh0e1/wire/interface/database"
	"github.com/sh0e1/wire/usecase"
)

// Injectors from wire.go:

func InitializeTodoHandler(sqlHander database.SQLHandler) *TodoHandler {
	todoRepository := database.NewTodoRepository(sqlHander)
	todoInteractor := usecase.NewTodoInteractor(todoRepository)
	todoHandler := NewTodoHandler(todoInteractor)
	return todoHandler
}
