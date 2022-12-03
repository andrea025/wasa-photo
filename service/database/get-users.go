package database


func (db *appdbimpl) GetUsers() ([]UserShortInfo, error) {
	users := []UserShortInfo{}

	sqlStmt := `SELECT * FROM User LIMIT 50;`
	rows, err := db.c.Query(sqlStmt)
	if err != nil {
		return []UserShortInfo{}, err
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