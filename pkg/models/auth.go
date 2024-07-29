package models

import (
	// "fmt"

	// "strings"

	"fmt"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"gorm.io/gorm"
	// "github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func AddNewUser(user *types.RegisterUserPayload) error {
	db, err := config.DbConnection()
	if err != nil {
		return err
	}
	tx := db.Create(&types.User{
		Username: user.Username, Email: user.Email, Password: user.Password})
	if tx.Error != nil {
		return err
	}
	err = createFirstUserAdmin(db)
	if err != nil {
		return err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}
func GetUserbyUserName(username string) (*types.User, error) {
	db, err := config.DbConnection()
	if err != nil {
		return nil, err
	}

	var user types.User
	tx := db.Model(&types.User{}).Where("username = ?", username).Scan(&user)

	if tx.Error != nil {
		return nil, err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func IsAdmin(userId int) (bool, error) {
	db, err := config.DbConnection()
	if err != nil {
		return false, err
	}
	var isAdmin bool
	tx := db.Table("users").Select("is_admin").Where("id = ?", userId).Scan(&isAdmin)
	if tx.Error != nil {
		return false, err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return false, err
	}
	return isAdmin, nil
}

func createFirstUserAdmin(db *gorm.DB) error {
	result := db.Model(&types.User{}).Where("id = 1 ").Update("is_admin", 1)
	if result.Error != nil {
		return result.Error
	}
	err := config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}
func GetAdminRequest() ([]types.AdminRequest, error) {
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	var adminRequests []types.AdminRequest
	rows, err := db.Model(&types.User{}).Select("id", "username", "email").Where("admin_request = TRUE").Rows()
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
	err = config.CloseConnection(db)
	if err != nil {
		return nil, err
	}
	return adminRequests, nil
}

func AdminRequest(userId int) error {
	db, err := config.DbConnection()
	if err != nil {
		return err
	}
	result := db.Model(&types.User{}).Where("id=?", userId).Update("admin_request", bool(true))
	if result.Error != nil {
		return result.Error
	}

	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}
func AdminRequestSent(userId int) (bool, error) {
	db, err := config.DbConnection()
	if err != nil {
		return false, err
	}
	type AdminRequestSent struct {
		Id      int
		request bool
	}
	payload := &AdminRequestSent{}
	tx := db.Table("users").Select("id", "admin_request").Where("id=?", userId).Scan(&payload)
	if tx.Error != nil {
		return false, err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return false, err
	}
	return payload.request, nil
}

// func ApproveAdmin(userIds []string) error {
// 	var keyString string
// 	for _, key := range userIds {
// 		keyString += key
// 		keyString += ","
// 	}
// 	keyString = strings.Trim(keyString, ",")
// 	query := fmt.Sprintf(`UPDATE user SET adminRequest = FALSE, isAdmin = TRUE WHERE id IN (%s)`, keyString)
// 	db, err := config.DbConnection()
// 	if err != nil {
// 		return err
// 	}
// 	result, err := db.Exec(query)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("error Updating")
// 	}
// 	err = config.CloseConnection(db)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func DenyAdminRequest(userId int) error {
// 	query := `UPDATE user SET adminRequest= false WHERE id= ?`
// 	db, err := config.DbConnection()
// 	if err != nil {
// 		return err
// 	}
// 	result, err := db.Exec(query, userId)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("error Updating")
// 	}
// 	err = config.CloseConnection(db)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
