package main

import (
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/cli"
	"github.com/refpiirmansah/pairproject_manufactures/configuration"
	"github.com/refpiirmansah/pairproject_manufactures/handler"
)

func main() {
	db, err := configuration.ConnectDB()
	if err != nil {
		fmt.Println("Koneksi database gagal:", err)
		return
	}
	defer db.Close()

	fmt.Println("Koneksi MySQL berhasil!")

	for {
		cli.ShowMenu()
		pilihan := cli.GetMenuChoice()

		switch pilihan {
		case 1:
			email, password := cli.GetLoginInput()

			user, err := handler.Login(db, email, password)

			if err != nil {
				fmt.Println("Login gagal:", err)
				continue
			}

			fmt.Println("Login berhasil!")
			fmt.Println("Selamat datang,", user.Email)
		case 0:
			fmt.Println("Terima kasih sudah menggunakan BuildTech PC")
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
