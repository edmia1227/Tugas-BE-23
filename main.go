package main

import (
	"Tugas-BE-23/configs"
	"Tugas-BE-23/internal/controller"
	"Tugas-BE-23/internal/model"
	"fmt"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	connection.AutoMigrate(&model.User{}, &model.Todo{})

	var inputMenu int
	um := model.NewUserModel(connection)
	uc := controller.NewUserController(um)

	tu := model.NewTodoModel(connection)
	tc := controller.NewTodoController(tu)

	// am := model.NewActivityModel(connection)
	// ac := controller.NewActivityController(am)

	for inputMenu != 9 {
		fmt.Println("Pilih menu")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)
		if inputMenu == 1 {
			var isLogin = true
			var inputMenu2 int
			data, err := uc.Login()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				return
			}

			for isLogin {
				fmt.Println("Selamat datang ", data.Name, ",")
				fmt.Println("Pilih menu")
				fmt.Println("1. Tambah Kegiatan")
				fmt.Println("2. Update Kegiatan")
				fmt.Println("3. Hapus Kegiatan")
				fmt.Println("4. Tampilkan daftar kegiatan")
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
				} else if inputMenu2 == 1 {
					_, err := tc.AddTodo(data.ID)
					if err != nil {
						fmt.Println("error ketika menambahkan aktivitas")
						return
					}
					fmt.Println("berhasil menambahkan aktivitas")
				}
			}

		} else if inputMenu == 2 {
			uc.Register()
		}
	}

	fmt.Println("terima kasih")

}
