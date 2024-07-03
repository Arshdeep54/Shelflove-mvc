package models

import (
	"fmt"
	"strconv"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/config"
	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func GetIssue(bookId string, userId int) (*types.Issue, error) {
	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		return nil, err
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
		err := rows.Scan(&request.Issue.Id, &request.Username, &request.Issue.Book_id, &request.Book.Title, &request.Book.Quantity, &request.Book.Author, &request.Issue.Issue_date, &request.Issue.Expected_return_date, &request.Issue.ReturnRequested, &request.Issue.IssueRequested)
		if err != nil {
			return nil, nil, fmt.Errorf("error scanning book data: %w", err)
		}
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
