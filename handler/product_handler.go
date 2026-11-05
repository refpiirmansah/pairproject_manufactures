package handler

import (
	"database/sql"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func GetAllProducts(db *sql.DB) ([]entity.Product, error) {
	query := `SELECT id, name, description, price, stock FROM products`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Stock,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func CreateProduct(db *sql.DB, name string, description string, price float64, stock int) error {
	query := `INSERT INTO products (name, description, price, stock) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, name, description, price, stock)
	return err
}

func UpdateProduct(db *sql.DB, id int, name string, description string, price float64, stock int) error {
	query := `UPDATE products SET name=?, description=?, price=?, stock=? WHERE id=?`
	_, err := db.Exec(query, name, description, price, stock, id)
	return err
}

func DeleteProduct(db *sql.DB, id int) error {
	query := `DELETE FROM products WHERE id=?`
	_, err := db.Exec(query, id)
	return err
}
