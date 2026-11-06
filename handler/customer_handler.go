package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func GetAllCustomers(db *sql.DB) ([]entity.Customer, error) {
	query := `
        SELECT id, name, email, phone, address
        FROM customers
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer

	for rows.Next() {
		var customer entity.Customer

		err := rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Email,
			&customer.Phone,
			&customer.Address,
		)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func CreateCustomer(db *sql.DB, customer entity.Customer) error {
	if customer.Name == "" {
		return fmt.Errorf("nama pelanggan wajib diisi")
	}

	if customer.Email == "" {
		return fmt.Errorf("email pelanggan wajib diisi")
	}

	query := `
        INSERT INTO customers (name, email, phone, address)
        VALUES (?, ?, ?, ?)
    `

	_, err := db.Exec(
		query,
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.Address,
	)

	return err
}

func UpdateCustomer(db *sql.DB, customer entity.Customer) error {
	if customer.ID <= 0 {
		return fmt.Errorf("ID pelanggan tidak valid")
	}

	if customer.Name == "" || customer.Email == "" {
		return fmt.Errorf("nama dan email wajib diisi")
	}

	query := `
        UPDATE customers
        SET name = ?, email = ?, phone = ?, address = ?
        WHERE id = ?
    `

	result, err := db.Exec(
		query,
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.Address,
		customer.ID,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("pelanggan dengan ID %d tidak ditemukan atau tidak berubah", customer.ID)
	}

	return nil
}

func DeleteCustomer(db *sql.DB, id int) error {
	if id <= 0 {
		return fmt.Errorf("ID pelanggan tidak valid")
	}

	query := `DELETE FROM customers WHERE id = ?`

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("pelanggan dengan ID %d tidak ditemukan", id)
	}

	return nil
}
