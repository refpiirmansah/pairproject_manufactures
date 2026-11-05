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

	// Buat user, jangan sampe di buka comment nya. nanti create user terus menerus
	//err = handler.CreateUser(db, "admin1@gmail.com", "abc")
	//if err != nil {
	//	fmt.Println("Gagal membuat akun pengujian:", err)
	//} else {
	//	fmt.Println("Akun pengujian berhasil dibuat")
	//}

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

			for {
				cli.ShowDashboard()
				pilihanDashboard := cli.GetMenuChoice()

				switch pilihanDashboard {
				case 1:
					// ------Memanggil sub-menu CRUD Komponen-----
					cli.ComponentMenu(db)
				case 2:
					// ------Memanggil sub-menu CRUD Produk------
					cli.ProductMenu(db)
				case 0:
					fmt.Println("Logout berhasil!")
					break
				default:
					fmt.Println("Fitur belum tersedia")
				}

				if pilihanDashboard == 0 {
					break
				}
			}
		case 0:
			fmt.Println("Terima kasih sudah menggunakan BuildTech PC")
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
