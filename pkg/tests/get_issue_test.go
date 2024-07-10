package tests

import (
	"database/sql"
	"strconv"

	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
)

type PayloadData struct {
	bookId string
	userId int
	err    error
}

var payloads = []PayloadData{
	{
		bookId: "",
		userId: 1,
		err:    strconv.ErrSyntax,
	},
	{
		bookId: "1",
		userId: 0,
		err:    sql.ErrNoRows,
	},
}

func TestGetIssue(t *testing.T) {
	
	config.DbPath = true
	for _, value := range payloads {
		_, err := models.GetIssue(value.bookId, value.userId)
		if err != nil {
			if err != value.err {
				t.Fatal(err.Error(), value.err)
			}
		}
	}
}
