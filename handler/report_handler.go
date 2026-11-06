package handler

import (
	"database/sql"
	"fmt"
)

func GetSalesReport(db *sql.DB) error {
	query := `
        SELECT
            COUNT(DISTINCT o.id),
            COALESCE(SUM(oi.quantity), 0),
            COALESCE(SUM(oi.quantity * oi.unit_price), 0)
        FROM orders o
        JOIN order_items oi
            ON oi.order_id = o.id
        WHERE o.status = 'completed'
    `

	var totalOrders int
	var totalPC int
	var totalSales float64

	err := db.QueryRow(query).Scan(
		&totalOrders,
		&totalPC,
		&totalSales,
	)
	if err != nil {
		return err
	}

	fmt.Println("\n=== LAPORAN PENJUALAN ===")
	fmt.Printf("Total Pesanan Selesai : %d\n", totalOrders)
	fmt.Printf("Total PC Terjual      : %d\n", totalPC)
	fmt.Printf("Total Nilai Penjualan : Rp%.2f\n", totalSales)

	return nil
}

func GetAssemblyReport(db *sql.DB) error {
	query := `
        SELECT
            COUNT(*),
            COALESCE(SUM(
                CASE WHEN status = 'processing' THEN 1 ELSE 0 END
            ), 0),
            COALESCE(SUM(
                CASE WHEN status = 'completed' THEN 1 ELSE 0 END
            ), 0),
            COALESCE(SUM(
                CASE WHEN status = 'completed' THEN quantity ELSE 0 END
            ), 0)
        FROM assemblies
    `

	var totalAssemblies int
	var processing int
	var completed int
	var totalPCCompleted int

	err := db.QueryRow(query).Scan(
		&totalAssemblies,
		&processing,
		&completed,
		&totalPCCompleted,
	)
	if err != nil {
		return err
	}

	fmt.Println("\n=== LAPORAN PERAKITAN ===")
	fmt.Printf("Total Perakitan       : %d\n", totalAssemblies)
	fmt.Printf("Sedang Diproses       : %d\n", processing)
	fmt.Printf("Selesai               : %d\n", completed)
	fmt.Printf("Total PC Selesai      : %d\n", totalPCCompleted)

	return nil
}

func GetComponentStockReport(db *sql.DB) error {
	query := `
        SELECT id, name, stock
        FROM components
        ORDER BY stock ASC, id ASC
    `

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("\n=== LAPORAN STOK KOMPONEN ===")

	found := false

	for rows.Next() {
		var componentID int
		var componentName string
		var stock int

		err := rows.Scan(
			&componentID,
			&componentName,
			&stock,
		)
		if err != nil {
			return err
		}

		status := "AMAN"

		if stock == 0 {
			status = "HABIS"
		} else if stock <= 5 {
			status = "MENIPIS"
		}

		fmt.Printf(
			"ID: %d | Komponen: %s | Stok: %d | Status: %s\n",
			componentID,
			componentName,
			stock,
			status,
		)

		found = true
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if !found {
		fmt.Println("Belum ada data komponen")
	}

	return nil
}
