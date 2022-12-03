package api

import (
	"github.com/julienschmidt/httprouter"
	"wasa-photo.uniroma1.it/wasa-photo/service/api/reqcontext"
	"wasa-photo.uniroma1.it/wasa-photo/service/database"
	"net/http"
	"fmt"
	"crypto/md5"
	"strings"
	"errors"
	"encoding/json"
	"time"
)


type Comment struct {
	Id string	`json:"id"`
	Photo string 	`json:"photo"`
	User string	`json:"user"`
	CreatedDatetime string	`json:"created_at"`
	Text string	`json:"text"`
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.Id = comment.Id
	c.Photo = comment.Photo
	c.User = comment.User
	c.CreatedDatetime = comment.CreatedDatetime
	c.Text = comment.Text
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		Id: c.Id,
		Photo: c.Photo,
		User: c.User,
		CreatedDatetime: c.CreatedDatetime,
		Text: c.Text,
	}
}

type CommentText struct {
	Text string `json:"text"`
}

func (ct *CommentText) isNotValid() bool {
	return len(ct.Text) < 3 || len(ct.Text) > 300
} 


func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var ct CommentText
	err := json.NewDecoder(r.Body).Decode(&ct)
	if err != nil {
		// The body was not a parseable JSON, reject it
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if ct.isNotValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment Comment
	pid := ps.ByName("photo_id")
	uid := strings.Split(r.Header.Get("Authorization"), "Bearer ")[1]
	creation := (time.Now()).Format(time.RFC3339)
	creation_datetime := creation[0:10] + " " + creation[11:19]		// correct format
	cid := fmt.Sprintf("%x", md5.Sum([]byte(pid + uid + creation)))
	comment = Comment{Id: cid, Photo: pid, User: uid, Text: ct.Text, CreatedDatetime: creation_datetime}

	dbComment, err := rt.db.CommentPhoto(comment.ToDatabase());
	if errors.Is(err, database.ErrBanned) {
		w.WriteHeader(http.StatusForbidden)
		return
	} else if errors.Is(err, database.ErrPhotoDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// some internal problem with the database
		ctx.Logger.WithError(err).Error("can't create comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comment.FromDatabase(dbComment)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
