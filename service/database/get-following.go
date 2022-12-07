package database

func (db *appdbimpl) GetFollowing(id string) ([]UserShortInfo, error) {
	users := []UserShortInfo{}

	exists, err := db.CheckUser(id)
	if err != nil {
		return []UserShortInfo{}, err
	} else if !exists {
		return []UserShortInfo{}, ErrUserDoesNotExist
	}

	sqlStmt := `SELECT user_followed, username FROM User, Following WHERE user_following == ? AND user_followed == id LIMIT 50;`
	rows, er := db.c.Query(sqlStmt, id)
	if er != nil {
		return []UserShortInfo{}, er
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var user UserShortInfo
		err = rows.Scan(&user.Id, &user.Username)
		if err != nil {
			return []UserShortInfo{}, err
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return []UserShortInfo{}, err
	}

	return users, nil
}