package user

import (
	"example.com/go-blog/db"
	"example.com/go-blog/shared"
)

func (u NewUser) Save() error {
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