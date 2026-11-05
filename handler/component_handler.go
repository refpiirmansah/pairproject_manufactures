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

func GetComponentByID(db *sql.DB, id int) (entity.Component, error) {
	var component entity.Component

	query := `SELECT id, name, category, stock, price FROM components WHERE id=?`

	err := db.QueryRow(query, id).Scan(
		&component.ID,
		&component.Name,
		&component.Category,
		&component.Stock,
		&component.Price,
	)

	if err != nil {
		return entity.Component{}, err
	}

	return component, nil
}

func UpdateComponent(db *sql.DB, component entity.Component) error {
	if component.ID < 0 {
		return fmt.Errorf("ID komponen tidak valid")
	}

	if component.Name == "" || component.Category == "" {
		return fmt.Errorf("Nama dan kategori wajib diisi")
	}

	if component.Stock < 0 || component.Stock < 0 {
		return fmt.Errorf("Stok dan harga tidak boleh negatif")
	}

	query := `
		UPDATE components 
		SET name = ?, category = ?, stock = ?, price = ?
		WHERE id = ?
	`

	result, err := db.Exec(
		query,
		component.Name,
		component.Category,
		component.Stock,
		component.Price,
		component.ID,
	)

	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("Komponen dengan ID %d tidak ditemukan", component.ID)
	}

	return nil
}
