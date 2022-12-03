package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) UncommentPhoto(photo_id string, comment_id string, req_id string) error {
	var id string
	err := db.c.QueryRow("SELECT id FROM Photo WHERE id == ?;", photo_id).Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrPhotoDoesNotExist
	} else if err != nil {
		return err
	} 

	var author string
	err = db.c.QueryRow("SELECT user FROM Comment WHERE id == ?;", comment_id).Scan(&author)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrCommentDoesNotExist
	} else if err != nil {
		return err
	}

	if author != req_id {
		return ErrForbidden
	}


	sqlStmt := `DELETE FROM Comment WHERE id=?;`
	_, err = db.c.Exec(sqlStmt, comment_id)
	return err
}
