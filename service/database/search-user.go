package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SearchUser(username string) (UserShortInfo, error) {
	var user UserShortInfo
	err := db.c.QueryRow("SELECT * FROM User WHERE username == ?;", username).Scan(&user.Id, &user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return user, ErrUserDoesNotExist
	} else if err != nil {
		return user, err
	}

	return user, nil
}
