package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func AddProductComponent(db *sql.DB, pc entity.ProductComponent) error {
	if pc.ProductID <= 0 || pc.ComponentID <= 0 {
		return fmt.Errorf("ID produk dan komponen tidak valid")
	}

	if pc.Quantity <= 0 {
		return fmt.Errorf("jumlah komponen harus lebih dari 0")
	}

	query := `
        INSERT INTO product_components
            (product_id, component_id, quantity)
        VALUES (?, ?, ?)
    `

	_, err := db.Exec(
		query,
		pc.ProductID,
		pc.ComponentID,
		pc.Quantity,
	)

	return err
}

func GetProductComponent(db *sql.DB, productID int) error {
	if productID <= 0 {
		return fmt.Errorf("ID product tidak valid")
	}

	query := `
		SELECT c.name, pc.quantity
		FROM product_components pc
		JOIN components c ON c.id = pc.component_id
		WHERE pc.product_id = ?
	`

	rows, err := db.Query(query, productID)
	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Println("\n=== RESEP RAKITAN ===")

	found := false

	for rows.Next() {
		var componentName string
		var quantity int

		err = rows.Scan(&componentName, &quantity)
		if err != nil {
			return err
		}

		fmt.Printf("%s | Jumlah: %d\n", componentName, quantity)
		found = true
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if !found {
		fmt.Println("Belum ada resep rakitan untuk produk ini")
	}

	return nil
}
