package controller

import (
	"Tugas-BE-23/internal/model"
	"fmt"
)

type ActivityController struct {
	model *model.ActivityModel
}

func NewActivityController(m *model.ActivityModel) *ActivityController {
	return &ActivityController{
		model: m,
	}
}

func (ac *ActivityController) AddActivity(id int) (bool, error) {
	var newData model.Activity
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.ActivityDate)
	newData.ID = id
	_, err := ac.model.AddActivity(newData)
	if err != nil {
		return false, err
	}
	return true, nil
}
