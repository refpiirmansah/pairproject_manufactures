package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func GetAllComponents(db *sql.DB) ([]entity.Component, error) {
	query := `SELECT id, name, category, stock, price FROM components`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var components []entity.Component

	for rows.Next() {
		var component entity.Component

		err := rows.Scan(
			&component.ID,
			&component.Name,
			&component.Category,
			&component.Stock,
			&component.Price,
		)

		if err != nil {
			return nil, err
		}

		components = append(components, component)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return components, nil
}

func CreateComponent(db *sql.DB, component entity.Component) error {
	if component.Name == "" || component.Category == "" {
		return fmt.Errorf("Nama dan kategori wajib diisi")
	}

	if component.Stock < 0 {
		return fmt.Errorf("Stok tidak boleh negatif")
	}

	if component.Price < 0 {
		return fmt.Errorf("Harga tidak boleh negatif")
	}

	query := `
		INSERT INTO components (name, category, stock, price) Values (?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		component.Name,
		component.Category,
		component.Stock,
		component.Price,
	)

	return err
}
