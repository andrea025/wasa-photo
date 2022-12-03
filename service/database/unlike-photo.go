package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) UnlikePhoto(photo_id string, user_id string) error {
	var id string
	err := db.c.QueryRow("SELECT id FROM Photo WHERE id == ?;", photo_id).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrPhotoDoesNotExist
	} else if err != nil {
		return err
	} 

	sqlStmt := `DELETE FROM Like WHERE photo == ? AND user == ?;`
	_, err = db.c.Exec(sqlStmt, photo_id, user_id)
	return err
}
