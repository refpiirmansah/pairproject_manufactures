package main

import (
	"fmt"
	"strings"

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
					components, err := handler.GetAllComponents(db)

					if err != nil {
						fmt.Println("Gagal mengambil data komponen:", err)
						continue
					}

					fmt.Println("\n=== DAFTAR KOMPONEN ===")

					if len(components) == 0 {
						fmt.Println("Belum ada data komponen")
						continue
					}

					fmt.Printf(
						"%-5s | %-20s | %-15s | %-6s | %-12s\n",
						"ID", "Nama", "Kategori", "Stok", "Harga",
					)

					fmt.Println(strings.Repeat("-", 72))

					for _, component := range components {
						fmt.Printf(
							"%-5d | %-20s | %-15s | %-6d | Rp%10.2f\n",
							component.ID,
							component.Name,
							component.Category,
							component.Stock,
							component.Price,
						)
					}

				case 2:
					products, err := handler.GetAllProducts(db)
					if err != nil {
						fmt.Println("Gagal mengambil data produk:", err)
					}

					if len(products) == 0 {
						fmt.Println("Belum ada data produk")
						continue
					}

					fmt.Println("\n=== DAFTAR PRODUK ===")
					fmt.Println(strings.Repeat("-", 85))
					fmt.Printf(
						"%-5s | %-20s | %-30s | %-6s | %-12s\n",
						"ID", "Nama", "Deskripsi", "Stok", "Harga",
					)
					for _, product := range products {
						fmt.Printf(
							"%-5d | %-20s | %-30s | %-6d | Rp%10.2f\n",
							product.ID,
							product.Name,
							product.Description,
							product.Stock,
							product.Price,
						)
					}

				case 7:
					component := cli.GetComponentInput()

					err := handler.CreateComponent(db, component)
					if err != nil {
						fmt.Println("Gagal menambahkan komponen:", err)
						continue
					}

					fmt.Println("komponen berhasil ditambahkan!")

				case 8:
					component := cli.GetUpdateComponentInput()

					err := handler.UpdateComponent(db, component)
					if err != nil {
						fmt.Println("Gagal mengubah komponen:", err)
						continue
					}
					fmt.Println("komponen berhasil diperbahui!")

				case 9:
					var id int
					var konfirmasi string

					fmt.Println("\n=== HAPUS KOMPONEN ===")
					fmt.Print("Masukkan ID komponen: ")
					_, err := fmt.Scanln(&id)

					if err != nil || id <= 0 {
						fmt.Println("ID komponen tidak valid")
						continue
					}

					fmt.Printf("Yakin ingin menghapus komponen ID %d? (y/n): ", id)
					fmt.Scanln(&konfirmasi)

					if konfirmasi != "y" {
						fmt.Println("Penghapusan dibatalkan")
						continue
					}

					err = handler.DeleteComponent(db, id)
					if err != nil {
						fmt.Println("Gagal menghapus komponen:", err)
						continue
					}

					fmt.Println("Komponen berhasil dihapus!")

				case 10:
					product := cli.GetProductInput()
					err := handler.CreateProduct(db, product)
					if err != nil {
						fmt.Println("Gagal menambahkan product:", err)
						continue
					}

					fmt.Println("product berhasil ditambahkan!")
				case 11:
					product := cli.GetUpdateProdukInput()

					err := handler.UpdateProduct(db, product)
					if err != nil {
						fmt.Println("Gagal mengubah produk:", err)
						continue
					}
					fmt.Println("produk berhasil diperbahui!")

				case 12:
					var id int
					var konfirmasi string

					fmt.Println("\n=== HAPUS PRODUK ===")
					fmt.Print("Masukkan ID produk: ")

					_, err := fmt.Scanln(&id)
					if err != nil || id <= 0 {
						fmt.Println("ID produk tidak valid")
						continue
					}

					fmt.Printf("Yakin ingin menghapus produk ID %d? (y/n): ", id)
					fmt.Scanln(&konfirmasi)

					if konfirmasi != "y" {
						fmt.Println("Penghapusan dibatalkan")
						continue
					}

					err = handler.DeleteProduct(db, id)
					if err != nil {
						fmt.Println("Gagal menghapus produk:", err)
						continue
					}

					fmt.Println("Produk berhasil dihapus!")

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
