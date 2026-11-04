package handler

import (
	"database/sql"
	"fmt"

	"github.com/refpiirmansah/pairproject_manufactures/entity"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByEmail(db *sql.DB, email string) (entity.User, error) {
	var user entity.User

	query := `
		SELECT id, email, password_hash
		FROM users
		WHERE email = ?
	`

	err := db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CreateUser(db *sql.DB, email string, password string) error {
	hash, err := HashPassword(password)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users(email, password_hash)
		VALUES (?, ?)
	`

	_, err = db.Exec(query, email, hash)
	if err != nil {
		return err
	}

	return nil
}

func CheckPassword(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))

	return err == nil
}

func Login(db *sql.DB, email string, password string) (entity.User, error) {
	user, err := GetUserByEmail(db, email)
	if err != nil {
		return entity.User{}, fmt.Errorf("email atau password salah")
	}

	if !CheckPassword(password, user.PasswordHash) {
		return entity.User{}, fmt.Errorf("email atau password salah")
	}

	return user, nil
}
