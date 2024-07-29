package tests

import (
	"database/sql"

	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

type GetUserIssuesTestData struct {
	userId int
	err    error
}

var getUserIssuesTestData = []GetUserIssuesTestData{
	{
		userId: 1,
		err:    sql.ErrNoRows,
	},
	{
		userId: 0,
		err:    sql.ErrNoRows,
	},
}

func TestGetUserIssues(t *testing.T) {
	config.DbPath = true
	for _, value := range getUserIssuesTestData {
		_, err := models.GetUserIssues(value.userId)
		if err != nil {
			if err != value.err {
				t.Fatal(err.Error(), value.err)
			}
		}
	}
}
