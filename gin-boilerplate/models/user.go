package models

import (
	"database/sql"
)

// User is user schema in mysql
type User struct {
	UserID   int    `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	IsAdmin  int    `db:"is_admin" json:"is_admin"`
}

// Save user
func (user *User) Save(tx *sql.Tx) (err error) {
	insertQuery := `INSERT INTO users (
						email,
						password,
						is_admin)
					VALUES ($1, $2, $3)`
	if tx != nil {
		result, err := tx.Exec(insertQuery, user.Email, user.Password, user.IsAdmin)

		if err != nil {
			return err
		}

		userID64, _ := result.LastInsertId()
		user.UserID = int(userID64)

		return err
	}
	_, err = DB.Exec(insertQuery, user.Email, user.Password, user.IsAdmin)

	return
}
