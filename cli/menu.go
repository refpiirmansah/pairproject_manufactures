package cli

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/refpiirmansah/pairproject_manufactures/handler"
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
	fmt.Println("0. Logout")
}

func ComponentMenu(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Kelola Komponen ===")
		fmt.Println("1. Lihat Komponen")
		fmt.Println("2. Tambah Komponen")
		fmt.Println("3. Ubah Komponen")
		fmt.Println("4. Hapus Komponen")
		fmt.Println("5. Kembali ke Dashboard")
		fmt.Print("Pilih menu: ")

		inputChoice, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(inputChoice))

		switch choice {
		case 1:
			components, _ := handler.GetAllComponents(db)
			fmt.Println("\n--- Daftar Komponen ---")
			for _, c := range components {
				fmt.Printf("ID: %d | Nama: %s | Kategori: %s | Stok: %d | Harga: %.2f\n", c.ID, c.Name, c.Category, c.Stock, c.Price)
			}
		case 2:
			fmt.Print("Nama: ")
			nameInput, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameInput)

			fmt.Print("Kategori: ")
			catInput, _ := reader.ReadString('\n')
			category := strings.TrimSpace(catInput)

			fmt.Print("Stok: ")
			stockInput, _ := reader.ReadString('\n')
			stock, _ := strconv.Atoi(strings.TrimSpace(stockInput))

			fmt.Print("Harga: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			handler.CreateComponent(db, name, category, stock, price)
			fmt.Println("Berhasil: Komponen baru telah ditambahkan.")
		case 3:
			fmt.Print("Input ID Komponen yang diubah: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			fmt.Print("Nama Baru: ")
			nameInput, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameInput)

			fmt.Print("Kategori Baru: ")
			catInput, _ := reader.ReadString('\n')
			category := strings.TrimSpace(catInput)

			fmt.Print("Stok Baru: ")
			stockInput, _ := reader.ReadString('\n')
			stock, _ := strconv.Atoi(strings.TrimSpace(stockInput))

			fmt.Print("Harga Baru: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			handler.UpdateComponent(db, id, name, category, stock, price)
			fmt.Println("Berhasil: Data komponen telah diperbarui.")
		case 4:
			fmt.Print("Input ID Komponen yang dihapus: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			handler.DeleteComponent(db, id)
			fmt.Println("Berhasil: Komponen telah dihapus.")
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func ProductMenu(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Kelola Produk ===")
		fmt.Println("1. Lihat Produk")
		fmt.Println("2. Tambah Produk")
		fmt.Println("3. Ubah Produk")
		fmt.Println("4. Hapus Produk")
		fmt.Println("5. Kembali ke Dashboard")
		fmt.Print("Pilih menu: ")

		inputChoice, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(inputChoice))

		switch choice {
		case 1:
			products, _ := handler.GetAllProducts(db)
			fmt.Println("\n--- Daftar Produk ---")
			for _, p := range products {
				fmt.Printf("ID: %d | Nama: %s | Deskripsi: %s | Harga: %.2f | Stok: %d\n", p.ID, p.Name, p.Description, p.Price, p.Stock)
			}
		case 2:
			fmt.Print("Nama Produk: ")
			nameInput, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameInput)

			fmt.Print("Deskripsi: ")
			descInput, _ := reader.ReadString('\n')
			description := strings.TrimSpace(descInput)

			fmt.Print("Harga: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			fmt.Print("Stok: ")
			stockInput, _ := reader.ReadString('\n')
			stock, _ := strconv.Atoi(strings.TrimSpace(stockInput))

			handler.CreateProduct(db, name, description, price, stock)
			fmt.Println("Berhasil: Produk baru telah ditambahkan.")
		case 3:
			fmt.Print("Input ID Produk yang diubah: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			fmt.Print("Nama Baru: ")
			nameInput, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameInput)

			fmt.Print("Deskripsi Baru: ")
			descInput, _ := reader.ReadString('\n')
			description := strings.TrimSpace(descInput)

			fmt.Print("Harga Baru: ")
			priceInput, _ := reader.ReadString('\n')
			price, _ := strconv.ParseFloat(strings.TrimSpace(priceInput), 64)

			fmt.Print("Stok Baru: ")
			stockInput, _ := reader.ReadString('\n')
			stock, _ := strconv.Atoi(strings.TrimSpace(stockInput))

			handler.UpdateProduct(db, id, name, description, price, stock)
			fmt.Println("Berhasil: Data produk telah diperbarui.")
		case 4:
			fmt.Print("Input ID Produk yang dihapus: ")
			idInput, _ := reader.ReadString('\n')
			id, _ := strconv.Atoi(strings.TrimSpace(idInput))

			handler.DeleteProduct(db, id)
			fmt.Println("Berhasil: Produk telah dihapus.")
		case 5:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
