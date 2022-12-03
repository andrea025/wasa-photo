package api

import (
	"github.com/julienschmidt/httprouter"
	"wasa-photo.uniroma1.it/wasa-photo/service/api/reqcontext"
	"wasa-photo.uniroma1.it/wasa-photo/service/database"
	"net/http"
	"errors"
	"encoding/json"
)


func (rt *_router) getBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	ret := []UserShortInfo{}

	dbfollowing, err := rt.db.GetBanned(ps.ByName("user_id"))
	if errors.Is(err, database.ErrUserDoesNotExist) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		// In this case, we have an error on our side. Log the error (so we can be notified) and send a 500 to the user
		// Note: we are using the "logger" inside the "ctx" (context) because the scope of this issue is the request.
		ctx.Logger.WithError(err).Error("can't list users this user is following")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, dbuser := range dbfollowing {
		var user UserShortInfo
		user.FromDatabase(dbuser)
		ret = append(ret, user)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(UsersInfoListResponse{Data:ret})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
