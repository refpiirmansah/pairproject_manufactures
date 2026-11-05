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
