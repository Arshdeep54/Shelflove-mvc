package tests

import (
	"database/sql"

	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

type GetRequestedAllTestData struct {
	err error
}

var getRequestedAllTestData = []GetRequestedAllTestData{
	{
		err: sql.ErrNoRows,
	},
}

func TestGetRequestedAll(t *testing.T) {
	config.DbPath = true
	_, _, err := models.GetRequestedAll()
	if err != nil {
		if err != getRequestedAllTestData[0].err {
			t.Fatal(err.Error(), getRequestedAllTestData[0].err)
		}
	}

}
