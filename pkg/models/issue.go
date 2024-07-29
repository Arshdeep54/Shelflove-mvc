package models

import (
	"fmt"
	"strconv"
	// "strings"
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
	row := db.Table("issues").Select("id", "is_returned", "return_requested", "issue_requested").Where("book_id = ? AND user_id = ? and is_returned=false ", int(id), userId).Row()
	if row.Err() != nil {
		return nil, row.Err()
	}
	var issue types.Issue
	err = row.Scan(&issue.Id, &issue.IsReturned, &issue.ReturnRequested, &issue.IssueRequested)
	if err != nil {
		return nil, err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return nil, err
	}
	return &issue, nil
}

func GetRequestedAll() ([]types.IssueWithDetails, []types.IssueWithDetails, error) {
	db, err := config.DbConnection()
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	var requests []types.IssueWithDetails
	rows, err := db.Model(&types.Issue{}).Joins("Inner Join users on users.id=issues.user_id").Joins("Inner Join books on books.id=issues.book_id").Where("return_requested = 1 or issue_requested = 1").Select("issues.id", "users.username", "books.id", "books.title", "books.quantity", "books.author", "issues.issue_date", "issues.expected_return_date", "issues.return_requested", "issues.issue_requested").Rows()
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
	err = config.CloseConnection(db)
	if err != nil {
		return nil, nil, err
	}
	return requestedReturns, requestedIssues, nil
}

func ExistingIssueCount(userId int) (int, error) {
	var count int
	db, err := config.DbConnection()
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}
	row := db.Model(&types.Issue{}).Where("user_id = ? AND is_returned = FALSE", userId).Select("Count(*)").Row()
	err = row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error connecting to Db: %w", err)
	}
	err = config.CloseConnection(db)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func AddNewIssue(userID int, bookId int) error {
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	tx := db.Model(&types.Issue{}).Create(&types.Issue{
		User_id: int32(userID), Book_id: int32(bookId), IssueRequested: true,
	})
	if tx.Error != nil {
		return err
	}
	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}
func ReturnRequest(userID int, bookId int) error {
	db, err := config.DbConnection()
	if err != nil {
		return fmt.Errorf("error connecting to Db: %w", err)
	}
	result := db.Model(&types.Issue{}).Where("user_id = ? AND book_id = ? AND isReturned=0", userID, bookId).Update("returnRequested", "true")
	if err != nil {
		return err
	}
	if result.Error != nil {
		return result.Error
	}

	err = config.CloseConnection(db)
	if err != nil {
		return err
	}
	return nil
}
func GetUserIssues(userId int) ([]types.IssueWithDetails, error) {

	db, err := config.DbConnection()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Db: %w", err)
	}
	var requests []types.IssueWithDetails
	rows, err := db.Model(&types.User{}).Joins("Inner Join issues on users.id=issues.user_id").Joins("Inner Join books on issues.book_id=books.id").Where("users.id=?", userId).Select("issues.id AS issueId", "issues.book_id", "issues.issue_date", "issues.expected_return_date", "books.title", "books.author", "issues.is_returned", "issues.issue_requested", "issues.return_requested").Rows()
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
	err = config.CloseConnection(db)
	if err != nil {
		return nil, err
	}
	return requests, nil
}

// func UpdateIssue(issueIds []string, updateType string) error {
// 	now := time.Now()
// 	nextDate := now.AddDate(0, 0, 14)
// 	formattedTodayDate := now.Format("2006-01-02")
// 	formattedReturnDate := nextDate.Format("2006-01-02")
// 	var keyString string
// 	for _, key := range issueIds {
// 		keyString += key
// 		keyString += ","
// 	}
// 	keyString = strings.Trim(keyString, ",")
// 	var query string
// 	if updateType == utils.ISSUED {

// 		query = fmt.Sprintf(` UPDATE issue SET issueRequested = FALSE, isReturned = FALSE ,issue_date='%s' , expected_return_date='%s' WHERE id IN (%s)`, formattedTodayDate, formattedReturnDate, keyString)
// 	} else if updateType == utils.RETURNED {
// 		query = fmt.Sprintf(` UPDATE issue SET returnRequested = FALSE, isReturned = TRUE ,returned_date ='%s' WHERE id IN (%s) `, formattedTodayDate, keyString)
// 	}
// 	db, err := config.DbConnection()
// 	if err != nil {
// 		return fmt.Errorf("error connecting to Db: %w", err)
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
// 		return fmt.Errorf("no Updation")
// 	}
// 	err = config.CloseConnection(db)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func DenyIssueRequest(id int, denyType string) error {
// 	var query string
// 	if denyType == utils.ISSUED {
// 		query = `DELETE FROM issue WHERE id= ?`
// 	} else if denyType == utils.RETURNED {
// 		query = ` UPDATE issue SET returnRequested = FALSE WHERE id = ?;`
// 	}
// 	db, err := config.DbConnection()
// 	if err != nil {
// 		return fmt.Errorf("error connecting to Db: %w", err)
// 	}
// 	result, err := db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("no Updation")
// 	}
// 	err = config.CloseConnection(db)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

// func BalanceIssues(userIds []string) error {
// 	var issues []types.IssueWithDetails
// 	for _, value := range userIds {
// 		id, err := strconv.ParseInt(value, 10, 64)
// 		if err != nil {
// 			return fmt.Errorf("error parsing to int")
// 		}
// 		userIssues, err := GetUserIssues(int(id))
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return err
// 		}
// 		issues = append(issues, userIssues...)
// 	}
// 	var issueIds []string
// 	var selectedBooks = map[string]int{}
// 	for _, value := range issues {
// 		if value.Issue.IssueRequested {
// 			err := DenyIssueRequest(value.Issue.Id, utils.ISSUED)
// 			if err != nil {

// 				fmt.Println(err.Error())
// 				return err
// 			}
// 		} else if !value.Issue.IsReturned && !value.Issue.IssueRequested {
// 			issueIds = append(issueIds, strconv.Itoa(value.Issue.Id))
// 			selectedBooks[strconv.Itoa(int(value.Issue.Book_id))] += 1
// 		}
// 	}
// 	payload := types.RequestPayload{
// 		IssueIds:      issueIds,
// 		SelectedBooks: selectedBooks,
// 	}
// 	err := UpdatebooksQuantity(&payload, true)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return err
// 	}
// 	err = UpdateIssue(issueIds, utils.RETURNED)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}
// 	return nil
// }
