package cli

import "fmt"

// Tampilan submenu tambahan. ShowDashboard dan ShowComponentMenu
// tetap menggunakan fungsi yang sudah ada di menu.go.
func ShowProductMenu() {
	fmt.Println("\n=== KELOLA PRODUK ===")
	fmt.Println("1. Lihat Produk")
	fmt.Println("2. Tambah Produk")
	fmt.Println("3. Update Produk")
	fmt.Println("4. Hapus Produk")
	fmt.Println("0. Kembali ke Dashboard")
}

func ShowCustomerMenu() {
	fmt.Println("\n=== KELOLA PELANGGAN ===")
	fmt.Println("1. Lihat Pelanggan")
	fmt.Println("2. Tambah Pelanggan")
	fmt.Println("3. Update Pelanggan")
	fmt.Println("4. Hapus Pelanggan")
	fmt.Println("0. Kembali ke Dashboard")
}

func ShowRecipeMenu() {
	fmt.Println("\n=== KELOLA RESEP RAKITAN ===")
	fmt.Println("1. Lihat Resep")
	fmt.Println("2. Tambah Komponen ke Resep")
	fmt.Println("3. Update Resep")
	fmt.Println("4. Hapus Komponen dari Resep")
	fmt.Println("5. Cek Stok untuk Perakitan")
	fmt.Println("0. Kembali ke Dashboard")
}

func ShowOrderMenu() {
	fmt.Println("\n=== KELOLA PESANAN ===")
	fmt.Println("1. Lihat Pesanan")
	fmt.Println("2. Buat Pesanan")
	fmt.Println("3. Update Status Pesanan")
	fmt.Println("4. Hapus Pesanan")
	fmt.Println("5. Selesaikan Pesanan")
	fmt.Println("0. Kembali ke Dashboard")
}

func ShowAssemblyMenu() {
	fmt.Println("\n=== KELOLA PERAKITAN ===")
	fmt.Println("1. Lihat Perakitan")
	fmt.Println("2. Mulai Perakitan")
	fmt.Println("3. Selesaikan Perakitan")
	fmt.Println("0. Kembali ke Dashboard")
}

func ShowReportMenu() {
	fmt.Println("\n=== LAPORAN ===")
	fmt.Println("1. Tampilkan Semua Laporan")
	fmt.Println("0. Kembali ke Dashboard")
}
