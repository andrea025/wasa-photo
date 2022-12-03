package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) CheckUser(user_id string) (bool, error) {
	var uid, username string
	err := db.c.QueryRow("SELECT * FROM User WHERE id == ?;", user_id).Scan(&uid, &username)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		// considering just the error, false will be ignored
		return false, err 
	}
	return true, nil
}
