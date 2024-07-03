package types

import "time"

type Issue struct {
	Id                   int
	User_id              int
	Book_id              int
	Issue_date           *time.Time
	Expected_return_date *time.Time
	Returned_date        *time.Time
	IsReturned           bool `default:"false"`
	ReturnRequested      bool `default:"false"`
	IssueRequested       bool `default:"false"`
	Fine                 int  `default:"0"`
}

type IssueWithDetails struct {
	Issue    Issue
	Book     Book
	Username string
}
