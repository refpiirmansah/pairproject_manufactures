package handler

import (
	"database/sql"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func GetAllCustomers(db *sql.DB) ([]entity.Customer, error) {
	query := `SELECT id, name, email, phone, address FROM customers`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer

	for rows.Next() {
		var c entity.Customer
		err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Email,
			&c.Phone,
			&c.Address,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func CreateCustomer(db *sql.DB, name, email, phone, address string) error {
	query := `INSERT INTO customers (name, email, phone, address) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, name, email, phone, address)
	return err
}

func UpdateCustomer(db *sql.DB, id int, name, email, phone, address string) error {
	query := `UPDATE customers SET name=?, email=?, phone=?, address=? WHERE id=?`
	_, err := db.Exec(query, name, email, phone, address, id)
	return err
}

func DeleteCustomer(db *sql.DB, id int) error {
	query := `DELETE FROM customers WHERE id=?`
	_, err := db.Exec(query, id)
	return err
}
