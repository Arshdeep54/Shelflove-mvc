package tests

import (
	"database/sql"

	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

type GetAdminRequestsTestData struct {
	userId int
	err    error
}

var getAdminRequestsTestData = []GetAdminRequestsTestData{
	{
		userId: 1,
		err:    sql.ErrNoRows,
	},
	{
		userId: 0,
		err:    sql.ErrNoRows,
	},
}

func TestGetAdminRequest(t *testing.T) {
	config.DbPath = true
	for _, value := range getAdminRequestsTestData {
		_, err := models.GetUserIssues(value.userId)
		if err != nil {
			if err != value.err {
				t.Fatal(err.Error(), value.err)
			}
		}
	}
}
