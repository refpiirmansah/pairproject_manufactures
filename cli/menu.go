package cli

import (
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

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

func ShowDashboard() {
	fmt.Println("\n=== DASHBOARD BUILDTECH PC ===")
	fmt.Println("1. Kelola Komponen")
	fmt.Println("2. Kelola Produk")
	fmt.Println("3. Kelola Pelanggan")
	fmt.Println("4. Kelola Pesanan")
	fmt.Println("5. Kelola Perakitan")
	fmt.Println("6. Laporan")
	fmt.Println("7. Tambah Komponen")
	fmt.Println("8. Update Komponen")
	fmt.Println("0. Logout")
}

func GetComponentInput() entity.Component {
	var component entity.Component

	fmt.Println("\n=== TAMBAH KOMPONEN ===")

	fmt.Print("Nama: ")
	fmt.Scanln(&component.Name)

	fmt.Print("Kategori: ")
	fmt.Scanln(&component.Category)

	fmt.Print("Stok: ")
	fmt.Scanln(&component.Stock)

	fmt.Print("Harga: ")
	fmt.Scanln(&component.Price)

	return component
}

func GetUpdateComponentInput() entity.Component {
	var component entity.Component

	fmt.Println("\n=== UPDATE KOMPONEN ===")

	fmt.Print("ID Komponen: ")
	fmt.Scanln(&component.ID)

	fmt.Print("Nama baru: ")
	fmt.Scanln(&component.Name)

	fmt.Print("Kategori baru: ")
	fmt.Scanln(&component.Category)

	fmt.Print("Stok baru: ")
	fmt.Scanln(&component.Stock)

	fmt.Print("Harga baru: ")
	fmt.Scanln(&component.Price)

	return component
}

func ShowUpdateComponentMenu() {
	fmt.Println("\n=== PILIH DATA YANG INGIN DIUBAH ===")
	fmt.Println("1. Nama")
	fmt.Println("2. Kategori")
	fmt.Println("3. Stok")
	fmt.Println("4. Harga")
	fmt.Println("0. Batal")
}
