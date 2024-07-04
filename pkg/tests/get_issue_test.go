package tests

import (
	// "fmt"
	"database/sql"
	"strconv"

	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/joho/godotenv"
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
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal("Error loading env", err)
	}
	for _, value := range payloads {
		_, err := models.GetIssue(value.bookId, value.userId)
		if err != nil {
			if err != value.err {
				t.Fatal(err.Error(), value.err)
			}
		}
	}
}
