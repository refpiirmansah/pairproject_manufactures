package cli

import "fmt"

func ShowMenu() {
	fmt.Println("\n=== BUILDTECH PC ===")
	fmt.Println("1. Login")
	fmt.Println("0. Keluar")
}

func GetMenuChoice() int {
	var pilihan int

	fmt.Print("Pilih menu: ")
	fmt.Scanln(&pilihan)

	return pilihan
}

func GetLoginInput() (string, string) {
	var email string
	var password string

	fmt.Println("\n=== LOGIN BUILDTECH PC ===")

	fmt.Print("Email: ")
	fmt.Scanln(&email)

	fmt.Print("Password: ")
	fmt.Scanln(&password)

	return email, password
}
