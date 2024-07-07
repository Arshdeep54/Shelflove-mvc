package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/utils"
)

func GetIssue(bookId string, userId int) (*types.Issue, error) {
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		return nil, strconv.ErrSyntax
	}
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	query := `SELECT id, isReturned,returnRequested,issueRequested FROM issue WHERE book_id = ? AND user_id = ? and isReturned=false `
	row := db.QueryRow(query, int(id), userId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var issue types.Issue
	err = row.Scan(&issue.Id, &issue.IsReturned, &issue.ReturnRequested, &issue.IssueRequested)
	if err != nil {
		return nil, err
	}
	return &issue, nil
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
	return adminRequests, nil
}
func GetRequestedAll() ([]types.IssueWithDetails, []types.IssueWithDetails, error) {
	db, err := config.DbConnection()
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	var requests []types.IssueWithDetails
	query := `
        SELECT i.id AS issueId, u.username,b.id AS bookId, b.title AS bookTitle,b.quantity,b.author, i.issue_date, i.expected_return_date, i.returnRequested ,i.issueRequested
        FROM issue i
        INNER JOIN user u ON u.id = i.user_id
        INNER JOIN book b ON b.id = i.book_id
        WHERE returnRequested = 1 or issueRequested = 1
      `
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, fmt.Errorf("error querying admin request: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var request types.IssueWithDetails
		var (
			issue    *time.Time
			expected *time.Time
		)
		err := rows.Scan(&request.Issue.Id, &request.Username, &request.Book.Id, &request.Book.Title, &request.Book.Quantity, &request.Book.Author, &issue, &expected, &request.Issue.ReturnRequested, &request.Issue.IssueRequested)
		if err != nil {
			return nil, nil, fmt.Errorf("error scanning book data: %w", err)
		}
		if expected == nil {

			request.Issue.Expected_return_date = "NOT ISSUED"
		} else {
			request.Issue.Expected_return_date = expected.Format(utils.LAYOUT)

		}
		if issue == nil {

			request.Issue.Issue_date = "NOT ISSUED"
		} else {
			request.Issue.Issue_date = issue.Format(utils.LAYOUT)

		}
		request.IsIssued = !request.Issue.IssueRequested
		requests = append(
			requests, request)
	}
	var requestedReturns, requestedIssues []types.IssueWithDetails

	for _, result := range requests {
		if result.Issue.ReturnRequested {
			requestedReturns = append(requestedReturns, result)
		}
		if result.Issue.IssueRequested {
			requestedIssues = append(requestedIssues, result)
		}
	}
	return requestedReturns, requestedIssues, nil
}

func ExistingIssueCount(userId int) (int, error) {
	var count int
	query := `SELECT Count(*) FROM issue WHERE user_id = ? AND isReturned = FALSE`
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}
	row := db.QueryRow(query, userId)
	err = row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}
	return count, nil
}
func AddNewIssue(userID int, bookId int) error {
	query := `INSERT INTO issue (user_id, book_id ,issueRequested) VALUES (?,?,?);`
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	_, err = db.Exec(query, userID, bookId, true)
	if err != nil {
		return err
	}
	return nil
}
func ReturnRequest(userID int, bookId int) error {
	query := `UPDATE issue
        SET returnRequested = true
        WHERE user_id = ? AND book_id = ? AND isReturned=0`
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	result, err := db.Exec(query, userID, bookId)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("issue Not Found")
	}
	return nil
}
func GetUserIssues(userId int) ([]types.IssueWithDetails, error) {
	query := `
        SELECT i.id AS issueId, i.book_id,i.issue_date, i.expected_return_date, b.title, b.author,i.isReturned,i.issueRequested,i.returnRequested
        FROM user u
        INNER JOIN issue i ON u.id = i.user_id
        INNER JOIN book b ON i.book_id = b.id
        WHERE u.id = ? ;
      `
	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	var requests []types.IssueWithDetails
	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, fmt.Errorf("error querying admin request: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var request types.IssueWithDetails
		var (
			issue    *time.Time
			expected *time.Time
		)
		err := rows.Scan(&request.Issue.Id, &request.Issue.Book_id, &issue, &expected, &request.Book.Title, &request.Book.Author, &request.Issue.IsReturned, &request.Issue.IssueRequested, &request.Issue.ReturnRequested)
		if err != nil {
			return nil, fmt.Errorf("error scanning book data: %w", err)
		}
		if expected == nil {

			request.Issue.Expected_return_date = "NOT ISSUED"
		} else {
			request.Issue.Expected_return_date = expected.Format(utils.LAYOUT)

		}
		if issue == nil {

			request.Issue.Issue_date = "NOT ISSUED"
		} else {
			request.Issue.Issue_date = issue.Format(utils.LAYOUT)

		}
		requests = append(
			requests, request)
	}
	return requests, nil
}

func UpdateIssue(issueIds []string, updateType string) error {
	now := time.Now()
	nextDate := now.AddDate(0, 0, 14)
	formattedTodayDate := now.Format("2006-01-02")
	formattedReturnDate := nextDate.Format("2006-01-02")
	var keyString string
	for _, key := range issueIds {
		keyString += key
		keyString += ","
	}
	keyString = strings.Trim(keyString, ",")
	var query string
	if updateType == utils.ISSUED {

		query = fmt.Sprintf(` UPDATE issue SET issueRequested = FALSE, isReturned = FALSE ,issue_date='%s' , expected_return_date='%s' WHERE id IN (%s)`, formattedTodayDate, formattedReturnDate, keyString)
	} else if updateType == utils.RETURNED {
		query = fmt.Sprintf(` UPDATE issue SET returnRequested = FALSE, isReturned = TRUE ,returned_date ='%s' WHERE id IN (%s) `, formattedTodayDate, keyString)
	}
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	println(rowsAffected)
	if rowsAffected == 0 {
		return fmt.Errorf("no Updation")
	}
	return nil
}

func DenyIssueRequest(id int, denyType string) error {
	var query string
	if denyType == utils.ISSUED {
		query = `DELETE FROM issue WHERE id= ?`
	} else if denyType == utils.RETURNED {
		query = ` UPDATE issue SET returnRequested = FALSE WHERE id = ?;`
	}
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no Updation")
	}
	return nil

}
