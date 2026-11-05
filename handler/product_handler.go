package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func GetAllProducts(db *sql.DB) ([]entity.Product, error) {
	query := `
			SELECT id, name, description, price, stock 
			FROM products
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var product entity.Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(db *sql.DB, product entity.Product) error {
	if product.Name == "" {
		return fmt.Errorf("nama produk wajib diisi")
	}

	if product.Price < 0 {
		return fmt.Errorf("harga produk tidak boleh negatif")
	}

	if product.Stock < 0 {
		return fmt.Errorf("stok produk tidak boleh negatif")
	}

	query := `
		INSERT INTO products (name, description, price, stock)
		VALUES (?, ?, ?, ?)
	`

	_, err := db.Exec(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)

	return err
}

func UpdateProduct(db *sql.DB, product entity.Product) error {
	if product.Name == "" {
		return fmt.Errorf("nama produk wajib diisi")
	}

	if product.Price < 0 {
		return fmt.Errorf("harga produk tidak boleh negatif")
	}

	if product.Stock < 0 {
		return fmt.Errorf("stok produk tidak boleh negatif")
	}

	query := `
		UPDATE products
		SET name = ?, description = ?, price = ?, stock = ?
		WHERE id = ?
	`

	result, err := db.Exec(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.ID,
	)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("Product  dengan ID %d tidak ditemukan", product.ID)
	}

	return nil
}

func DeleteProduct(db *sql.DB, id int) error {
	if id <= 0 {
		return fmt.Errorf("ID produk tidak valid")
	}

	query := `DELETE FROM products WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("produk dengan ID %d tidak ditemukan", id)
	}

	return nil
}
