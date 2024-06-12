package controller

import (
	"Tugas-BE-23/internal/model"
	"errors"
	"fmt"
)

type UserController struct {
	model *model.UserModel
}

func NewUserController(m *model.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Login() (model.User, error) {
	var email, password string
	fmt.Print("Masukkan email ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan password ")
	fmt.Scanln(&password)
	result, err := uc.model.Login(email, password)
	if err != nil {
		return model.User{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}

func (uc *UserController) Register() (model.User, error) {
	var newData model.User
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newData.Name)
	fmt.Print("Masukkan Password ")
	fmt.Scanln(&newData.Password)
	fmt.Print("Masukkan Email ")
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan HP ")
	fmt.Scanln(&newData.Phone)
	result, err := uc.model.Register(newData)
	if err != nil && !result {
		return model.User{}, errors.New("terjadi masalah ketika registrasi")
	}
	return newData, nil
}
