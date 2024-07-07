package models

import (
	"database/sql"
)

type User struct {
	ID       string
	Email    string
	Username string
	Password string
	Admin    bool
}

func (user *User) Create(db *sql.DB) error {
	query := `INSERT INTO public.users (email, username, password, admin) VALUES ($1, $2, $3, $4) RETURNING id`
	return db.QueryRow(query, user.Email, user.Username, user.Password, user.Admin).Scan(&user.ID)
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, username, password, admin FROM public.users WHERE username = $1`
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Admin)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// TODO: Move methods to handler
