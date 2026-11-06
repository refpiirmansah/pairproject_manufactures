package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func CreateOrder(db *sql.DB, customerID int) (int64, error) {
	if customerID <= 0 {
		return 0, fmt.Errorf("ID pelanggan tidak valid")
	}

	query := `
		INSERT INTO orders (customer_id, order_date, status)
		VALUES (?, ?, ?)
	`

	result, err := db.Exec(query, customerID, "pending")
	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

func CreateOrderItem(db *sql.DB, item entity.OrderItem) error {
	if item.OrderID <= 0 || item.ProductID <= 0 {
		return fmt.Errorf("ID pesanan atau produk tidak valid")
	}

	if item.Quantity <= 0 {
		return fmt.Errorf("jumlah produk harus lebih dari 0")
	}

	if item.UnitPrice < 0 {
		return fmt.Errorf("harga produk tidak boleh negatif")
	}

	query := `
		INSERT INTO order_items (order_id, product_id, quantity, unit_price)
		VALUES (?, ?, ?, ?)
	`

	_, err := db.Exec(query, item.OrderID, item.ProductID, item.Quantity, item.UnitPrice)

	return err
}

func CreateOrderWithItem(
	db *sql.DB,
	customerID int,
	productID int,
	quantity int,
) (int64, error) {
	if customerID <= 0 || productID <= 0 || quantity <= 0 {
		return 0, fmt.Errorf("data pesanan tidak valid")
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Ambil harga produk langsung dari database
	var price float64

	err = tx.QueryRow(
		"SELECT price FROM products WHERE id = ?",
		productID,
	).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("gagal mengambil harga produk: %w", err)
	}

	// Simpan data utama pesanan
	result, err := tx.Exec(
		`INSERT INTO orders (customer_id, order_date, status)
         VALUES (?, NOW(), ?)`,
		customerID,
		"pending",
	)
	if err != nil {
		return 0, err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Simpan detail produk yang dipesan
	_, err = tx.Exec(
		`INSERT INTO order_items
         (order_id, product_id, quantity, unit_price)
         VALUES (?, ?, ?, ?)`,
		orderID,
		productID,
		quantity,
		price,
	)
	if err != nil {
		return 0, err
	}

	// Simpan kedua perubahan secara permanen
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return orderID, nil
}

func GetAllOrders(db *sql.DB) error {
	query := `
        SELECT
            o.id,
            c.name,
            p.name,
            oi.quantity,
            oi.unit_price,
            o.status
        FROM orders o
        JOIN customers c ON o.customer_id = c.id
        JOIN order_items oi ON o.id = oi.order_id
        JOIN products p ON oi.product_id = p.id
        ORDER BY o.id
    `
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== DAFTAR PESANAN ===")

	found := false

	for rows.Next() {
		var orderID int
		var customerName string
		var productName string
		var quantity int
		var unitPrice float64
		var status string

		err = rows.Scan(&orderID, &customerName, &productName, &quantity, &unitPrice)
		if err != nil {
			return err
		}

		fmt.Printf(
			"Order ID: %d | Pelanggan: %s | Produk: %s | Jumlah: %d | Harga: Rp%.2f | Status: %s\n",
			orderID,
			customerName,
			productName,
			quantity,
			unitPrice,
			status,
		)

		found = true
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if !found {
		fmt.Println("Belum ada pesanan")
	}

	return nil
}
