package handler

import (
	"database/sql"

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

func CreateComponent(db *sql.DB, name string, category string, stock int, price float64) error {
	query := `INSERT INTO components (name, category, stock, price) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, name, category, stock, price)
	return err
}

func UpdateComponent(db *sql.DB, id int, name string, category string, stock int, price float64) error {
	query := `UPDATE components SET name=?, category=?, stock=?, price=? WHERE id=?`
	_, err := db.Exec(query, name, category, stock, price, id)
	return err
}

func DeleteComponent(db *sql.DB, id int) error {
	query := `DELETE FROM components WHERE id=?`
	_, err := db.Exec(query, id)
	return err
}
