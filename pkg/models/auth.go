package models

import (
	"database/sql"
	"fmt"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func AddNewUser(user *types.RegisterUserPayload) error {
	query := `INSERT INTO user (username, email, password) VALUES (?, ?, ?)`
	db, err := config.DbConnection()
	if err != nil {
		return err
	}
	_, err = db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	err = createFirstUserAdmin(db)
	if err != nil {
		return err
	}
	return nil
}
func GetUserbyUserName(username string) (*types.User, error) {
	query := `SELECT * FROM user WHERE username = ?`
	db, err := config.DbConnection()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow(query, username)

	var user types.User
	err = row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.IsAdmin, &user.AdminRequest)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func createFirstUserAdmin(db *sql.DB) error {
	query := `UPDATE user SET isAdmin=true WHERE id=1`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
func AdminRequest(userId int) error {
	query := `UPDATE user SET adminRequest = TRUE WHERE id=?`
	db, err := config.DbConnection()
	if err != nil {
		return err
	}
	result, err := db.Exec(query, userId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("error Updating")
	}

	return nil
}
func AdminRequestSent(userId int) (bool, error) {
	query := `
    SELECT id,adminRequest
    FROM user
    WHERE id = ? 
  ;`
	db, err := config.DbConnection()
	if err != nil {
		return false, err
	}
	row := db.QueryRow(query, userId)
	var (
		id      int
		request bool
	)
	err = row.Scan(&id, &request)
	if err != nil {
		return false, err
	}
	return request, nil
}
