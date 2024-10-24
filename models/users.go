package models

import (
	"errors"

	"example.com/go-rest-api/db"
	"example.com/go-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
	SELECT id, password FROM users
	WHERE email = ?
	`

	var retrievedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(retrievedPassword, u.Password)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
