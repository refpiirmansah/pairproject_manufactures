package main

import (
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/cli"
	"github.com/refpiirmansah/pairproject_manufactures/configuration"
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
			fmt.Println("Menu Login dipilih")
		case 0:
			fmt.Println("Terima kasih sudah menggunakan BuildTech PC")
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}
