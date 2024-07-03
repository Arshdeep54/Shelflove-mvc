package models

import (
	"database/sql"

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
   err=row.Scan(&user.Id,&user.Username,&user.Email,&user.Password,&user.IsAdmin,&user.AdminRequest)
   if err!=nil{
      return nil,err
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
