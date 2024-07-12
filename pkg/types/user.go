package types

type User struct {
	Id           int32  `json:"id" gorm:"primaryKey"`
	Username     string `json:"username" gorm:"unique"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password"`
	IsAdmin      bool   `json:"isAdmin"`
	AdminRequest bool   `json:"adminRequest"`
}

type AdminRequest struct {
	Id       int32
	Username string
	Email    string `json:"email"`
}
type RegisterUserPayload struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

type JwtPayload struct {
	Id           int32  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"isAdmin"`
	AdminRequest bool   `json:"adminRequest"`
}
type RenderData struct {
	IsAdmin           bool
	IsLoggedIn        bool
	UserId            int
	Books             []Book
	Book              *Book
	ErrorMessage      string
	IssueRequested    bool
	IsIssued          bool
	IsReturnRequested bool
	Username          string
	Email             string
	RequestedReturns  []IssueWithDetails
	RequestedIssues   []IssueWithDetails
	RequestedAdmins   []AdminRequest
	AdminRequested    bool
	HomeActive        bool
	BooksActive       bool
}

type RequestPayload struct {
	IssueIds      []string       `json:"issueIds"`
	SelectedBooks map[string]int `json:"selectedBooks"`
}
