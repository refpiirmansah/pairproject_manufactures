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

func UpdateProductComponent(db *sql.DB, productID int, componentID int, quantity int) error {
	if productID <= 0 || componentID <= 0 {
		return fmt.Errorf("ID produk dan komponen tidak valid")
	}

	if quantity <= 0 {
		return fmt.Errorf("jumlah komponen harus lebih dari 0")
	}

	query := `
		UPDATE product_components
		SET quantity = ?
		WHERE product_id = ? AND component_id = ?
	`

	result, err := db.Exec(query, quantity, productID, componentID)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect == 0 {
		return fmt.Errorf("resep tidak ditemukan atau jumlah tidak berubah")
	}

	return nil
}

func DeleteProductComponent(db *sql.DB, productID int, componentID int) error {
	if productID <= 0 || componentID <= 0 {
		return fmt.Errorf("ID produk dan komponen tidak valid")
	}

	query := `
		DELETE FROM product_components
		WHERE product_id = ? AND component_id = ?
	`

	result, err := db.Exec(query, productID, componentID)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect == 0 {
		fmt.Errorf("komponen tidak ditemukan dalam resep produk")
	}

	return nil
}

func ValidateRecipeStock(db *sql.DB, productID, totalPC int) error {
	if productID <= 0 || totalPC <= 0 {
		return fmt.Errorf("ID produk dan jumlah PC harus lebih dari 0")
	}

	query := `
        SELECT c.name, c.stock, pc.quantity
        FROM product_components pc
        JOIN components c ON pc.component_id = c.id
        WHERE pc.product_id = ?
    `

	rows, err := db.Query(query, productID)
	if err != nil {
		return err
	}

	defer rows.Close()

	found := false
	for rows.Next() {
		found = true

		var name string
		var stock int
		var quantityPerPC int

		if err := rows.Scan(&name, &stock, &quantityPerPC); err != nil {
			return err
		}

		if !ValidateComponentStock(stock, quantityPerPC, totalPC) {
			return fmt.Errorf("stok %s tidak cukup untuk %d PC", name, totalPC)
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("produk ini belum memiliki resep rakitan")
	}

	return nil
}
