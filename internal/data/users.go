package data

import "database/sql"

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Version   string `json:"version"`
}

type UserModel struct {
	BD *sql.DB
}
