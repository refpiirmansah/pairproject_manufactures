package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
)

func GetOrderItemForAssembly(
	db *sql.DB,
	orderItemID int,
) (int, int, error) {
	if orderItemID <= 0 {
		return 0, 0, fmt.Errorf("ID detail pesanan tidak valid")
	}

	var productID int
	var quantity int
	var orderStatus string

	query := `
        SELECT
            oi.product_id,
            oi.quantity,
            o.status
        FROM order_items oi
        JOIN orders o ON oi.order_id = o.id
        WHERE oi.id = ?
    `

	err := db.QueryRow(query, orderItemID).Scan(
		&productID,
		&quantity,
		&orderStatus,
	)

	if err == sql.ErrNoRows {
		return 0, 0, fmt.Errorf("detail pesanan tidak ditemukan")
	}

	if err != nil {
		return 0, 0, err
	}

	if orderStatus != "processing" {
		return 0, 0, fmt.Errorf(
			"pesanan harus berstatus processing sebelum dirakit",
		)
	}

	return productID, quantity, nil
}

func ValidateAssemblyStock(db *sql.DB, orderItemID int) error {
	//Ambil produk dan jumlah PC dari pesanan
	productID, quantity, err := GetOrderItemForAssembly(db, orderItemID)
	if err != nil {
		return err
	}

	//periksa stock semua komponen resep PC
	err = ValidateRecipeStock(db, productID, quantity)
	if err != nil {
		return fmt.Errorf(
			"stok komponen tidak mencukupi: %w",
			err,
		)
	}

	return nil
}

func CreateAssembly(db *sql.DB, orderItemID int) (int64, error) {
	if orderItemID <= 0 {
		return 0, fmt.Errorf("ID detail pesanan tidak valid")
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	// Ambil detail pesanan dan kunci barisnya
	var productID int
	var quantity int
	var orderStatus string

	err = tx.QueryRow(
		`SELECT oi.product_id, oi.quantity, o.status
         FROM order_items oi
         JOIN orders o ON oi.order_id = o.id
         WHERE oi.id = ?
         FOR UPDATE`,
		orderItemID,
	).Scan(&productID, &quantity, &orderStatus)

	if err == sql.ErrNoRows {
		return 0, fmt.Errorf("detail pesanan tidak ditemukan")
	}
	if err != nil {
		return 0, err
	}

	if orderStatus != "processing" {
		return 0, fmt.Errorf(
			"pesanan harus berstatus processing",
		)
	}

	// Cegah perakitan ganda
	var exists bool

	err = tx.QueryRow(
		`SELECT EXISTS(
            SELECT 1 FROM assemblies
            WHERE order_item_id = ?
        )`,
		orderItemID,
	).Scan(&exists)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, fmt.Errorf(
			"perakitan untuk detail pesanan ini sudah ada",
		)
	}

	// Hitung seluruh kebutuhan komponen
	components, err := GetAssemblyComponentNeedsTx(
		tx,
		productID,
		quantity,
	)
	if err != nil {
		return 0, err
	}

	// Buat catatan perakitan
	result, err := tx.Exec(
		`INSERT INTO assemblies
            (order_item_id, quantity, start_date, end_date, status)
         VALUES (?, ?, NOW(), NULL, ?)`,
		orderItemID,
		quantity,
		"processing",
	)
	if err != nil {
		return 0, err
	}

	assemblyID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Kurangi stok dan catat komponen
	err = SaveAssemblyComponents(
		tx,
		assemblyID,
		components,
	)
	if err != nil {
		return 0, err
	}

	// Simpan semua perubahan
	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return assemblyID, nil
}

func CalculateComponentNeeds(quantityPerPC int, totalPC int) (int, error) {
	if quantityPerPC <= 0 {
		return 0, fmt.Errorf(
			"jumlah komponen per PC harus lebih dari 0",
		)
	}

	if totalPC <= 0 {
		return 0, fmt.Errorf(
			"jumlah PC harus lebih dari 0",
		)
	}

	totalNeeded := quantityPerPC * totalPC
	return totalNeeded, nil
}

func GetAssemblyComponentNeeds(db *sql.DB, productID int, totalPC int) ([]entity.AssemblyComponent, error) {
	if productID <= 0 || totalPC <= 0 {
		return nil, fmt.Errorf("data perakitan tidak valid")
	}

	query := `
		SELECT component_id, quantity
		FROM product_components
		WHERE product_id = ?
	`

	rows, err := db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []entity.AssemblyComponent

	for rows.Next() {
		var componentID int
		var quantityPerPC int

		err = rows.Scan(&componentID, &quantityPerPC)
		if err != nil {
			return nil, err
		}

		totalNeeded, err := CalculateComponentNeeds(quantityPerPC, totalPC)
		if err != nil {
			return nil, err
		}

		components = append(components, entity.AssemblyComponent{
			ComponentID:  componentID,
			QuantityUsed: totalNeeded,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(components) == 0 {
		return nil, fmt.Errorf("resep komponen produk belum tersedia")
	}

	return components, nil
}

func SaveAssemblyComponents(
	tx *sql.Tx,
	assemblyID int64,
	components []entity.AssemblyComponent,
) error {
	if assemblyID <= 0 || len(components) == 0 {
		return fmt.Errorf("data komponen perakitan tidak valid")
	}

	for _, component := range components {
		if component.ComponentID <= 0 || component.QuantityUsed <= 0 {
			return fmt.Errorf("data pemakaian komponen tidak valid")
		}

		// Kurangi stok hanya jika stok mencukupi
		result, err := tx.Exec(
			`UPDATE components
             SET stock = stock - ?
             WHERE id = ? AND stock >= ?`,
			component.QuantityUsed,
			component.ComponentID,
			component.QuantityUsed,
		)
		if err != nil {
			return err
		}

		affected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if affected == 0 {
			return fmt.Errorf(
				"stok komponen ID %d tidak mencukupi",
				component.ComponentID,
			)
		}

		// Catat jumlah komponen yang digunakan
		_, err = tx.Exec(
			`INSERT INTO assembly_components
                (assembly_id, component_id, quantity_used)
             VALUES (?, ?, ?)`,
			assemblyID,
			component.ComponentID,
			component.QuantityUsed,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAssemblyComponentNeedsTx(tx *sql.Tx, productID int, totalPC int) ([]entity.AssemblyComponent, error) {
	if productID <= 0 || totalPC <= 0 {
		return nil, fmt.Errorf("data perakitan tidak valid")
	}

	query := `
        SELECT component_id, quantity
        FROM product_components
        WHERE product_id = ?
        ORDER BY component_id
    `

	rows, err := tx.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []entity.AssemblyComponent
	for rows.Next() {
		var componentID int
		var quantityPerPC int

		err := rows.Scan(&componentID, &quantityPerPC)
		if err != nil {
			return nil, err
		}

		totalNeeded, err := CalculateComponentNeeds(quantityPerPC, totalPC)
		if err != nil {
			return nil, err
		}

		components = append(components, entity.AssemblyComponent{
			ComponentID:  componentID,
			QuantityUsed: totalNeeded,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(components) == 0 {
		return nil, fmt.Errorf(
			"resep komponen produk belum tersedia",
		)
	}

	return components, nil
}
