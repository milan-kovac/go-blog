package user

import (
	"errors"

	"example.com/go-blog/db"
	"example.com/go-blog/shared"
)

func (u UserData) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	
	hashedPassword, err := shared.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_ , err = stmt.Exec(u.Email, hashedPassword)


	return err
}



func (u *UserData) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)


	if err != nil {
		return err
	}

	passwordIsValid := shared.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid.")
	}

	return nil
}