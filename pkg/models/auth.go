package models

import (
	"database/sql"
	"fmt"
	"strings"

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
	db.Close()
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
	db.Close()
	return &user, nil
}
func IsAdmin(userId int) (bool, error) {
	query := `SELECT isAdmin FROM user WHERE id = ?`
	db, err := config.DbConnection()
	if err != nil {
		return false, err
	}
	row := db.QueryRow(query, userId)
	var isAdmin bool
	err = row.Scan(&isAdmin)
	if err != nil {
		return false, err
	}
	db.Close()
	return isAdmin, nil
}
func createFirstUserAdmin(db *sql.DB) error {
	query := `UPDATE user SET isAdmin=true WHERE id=1`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	db.Close()
	return nil
}
func GetAdminRequest() ([]types.AdminRequest, error) {
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	var adminRequests []types.AdminRequest
	query := `
      SELECT u.id AS userId, u.username, u.email
      FROM user u
      WHERE adminRequest = TRUE
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying admin request: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var adminRequest types.AdminRequest
		err := rows.Scan(&adminRequest.Id, &adminRequest.Username, &adminRequest.Email)
		if err != nil {
			return nil, fmt.Errorf("error scanning book data: %w", err)
		}
		adminRequests = append(adminRequests, adminRequest)
	}
	db.Close()
	return adminRequests, nil
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
	db.Close()
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
	db.Close()
	return request, nil
}

func ApproveAdmin(userIds []string) error {
	var keyString string
	for _, key := range userIds {
		keyString += key
		keyString += ","
	}
	keyString = strings.Trim(keyString, ",")
	query := fmt.Sprintf(`UPDATE user SET adminRequest = FALSE, isAdmin = TRUE WHERE id IN (%s)`, keyString)
	db, err := config.DbConnection()
	if err != nil {
		return err
	}
	result, err := db.Exec(query)
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
	db.Close()
	return nil
}
func DenyAdminRequest(userId int) error {
	query := `UPDATE user SET adminRequest= false WHERE id= ?`
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
	db.Close()

	return nil
}
