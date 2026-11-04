package main

import (
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/cli"
)

func main() {
	for {
		var pilihan int

		cli.ShowMenu()

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

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
