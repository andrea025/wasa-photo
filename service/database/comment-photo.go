package database

import (
	"database/sql"
	"errors"
)

type Comment struct {
	Id              string
	Photo           string
	User            string
	CreatedDatetime string
	Text            string
}

func (db *appdbimpl) CommentPhoto(comment Comment) (Comment, error) {
	var owner_id string
	err := db.c.QueryRow("SELECT owner FROM Photo WHERE id == ?;", comment.Photo).Scan(&owner_id)
	if errors.Is(err, sql.ErrNoRows) {
		return comment, ErrPhotoDoesNotExist
	} else if err != nil {
		return comment, err
	}

	var banned bool
	banned, err = db.CheckBan(owner_id, comment.User)
	if err != nil {
		return comment, err
	} else if banned {
		return comment, ErrBanned
	}

	sqlStmt := `INSERT INTO Comment (id, photo, user, text, created_at) VALUES (?, ?, ?, ?, ?)`
	_, err = db.c.Exec(sqlStmt, comment.Id, comment.Photo, comment.User, comment.Text, comment.CreatedDatetime)
	return comment, err
}
