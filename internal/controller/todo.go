package controller

import (
	"Tugas-BE-23/internal/model"
	"fmt"
)

type TodoController struct {
	model *model.TodoModel
}

func NewTodoController(m *model.TodoModel) *TodoController {
	return &TodoController{
		model: m,
	}
}

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData model.Todo
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = id
	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}
