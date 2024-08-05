package tests_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/models"
	"github.com/DATA-DOG/go-sqlmock"     // Import mock library
	"github.com/stretchr/testify/assert" // Import assertion library
	"github.com/stretchr/testify/require"
)

type Database interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Close() error
}

type sqlDatabase struct {
	db *sql.DB
}

func (s *sqlDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	return s.db.QueryRow(query, args...)
}

func TestGetIssue(t *testing.T) {
	t.Run("Issue Exists", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.Error(t, err)

		defer mockDB.Close()

		rows := mock.NewRows([]string{"id", "isReturned", "returnRequested", "issueRequested"}).
			AddRow(1, false, false, true)
		mock.ExpectQuery(`SELECT id, isReturned,returnRequested,issueRequested FROM issue WHERE book_id = \? AND user_id = \? and isReturned=false  `).
			WithArgs(1, 1).
			WillReturnRows(rows)

		bookId := "1"
		userId := 1
		mockDatabase := &sqlDatabase{db: mockDB}
		issue, err := models.GetIssue(mockDatabase.db, bookId, userId)
		require.NoError(t, err)
		assert.NotNil(t, issue)
		assert.Equal(t, int(1), issue.Id)
		assert.False(t, false, issue.IsReturned)
		assert.False(t, false, issue.ReturnRequested)
		assert.False(t, true, issue.IssueRequested)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("Issue Not Found", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		defer mockDB.Close()
		mock.ExpectQuery(`SELECT id, isReturned,returnRequested,issueRequested FROM issue WHERE book_id = \? AND user_id = \? and isReturned=false  `).WithArgs(2, 1).WillReturnError(fmt.Errorf("no rows in result set"))

		bookId := "2"
		userId := 1
		mockDatabase := &sqlDatabase{db: mockDB}
		issue, err := models.GetIssue(mockDatabase.db, bookId, userId)
		require.Error(t, err)
		require.Nil(t, issue)
		require.EqualError(t, err, "no rows in result set")
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
