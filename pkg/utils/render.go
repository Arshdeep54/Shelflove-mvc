package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Arshdeep54/Shelflove-mvc/pkg/types"
)

func UpdateHeaders(w http.ResponseWriter, Data *types.RenderData) error {
	isAdmin, err := strconv.ParseBool(w.Header().Get("IsAdmin"))
	if err != nil {
		return err
	}
	IsLoggedIn, err := strconv.ParseBool(w.Header().Get("IsLoggedIn"))
	if err != nil {
		return err
	}
	userId, err := strconv.ParseInt(w.Header().Get("userId"), 10, 64)
	if err != nil {
		return err
	}
	if len(strings.Split(w.Header().Get("issueRequested"), "")) > 0 {
		IssueRequested, err := strconv.ParseBool(w.Header().Get("issueRequested"))
		Data.IssueRequested = IssueRequested
		if err != nil {
			return err
		}
	}

	if len(strings.Split(w.Header().Get("isIssued"), "")) > 0 {
		isIssued, err := strconv.ParseBool(w.Header().Get("isIssued"))
		Data.IsIssued = isIssued
		fmt.Println(Data.IsIssued, w.Header().Get("isIssued"), w.Header().Get("issueRequested"), w.Header().Get("isReturnRequested"))
		if err != nil {
			return err
		}
	}
	if len(strings.Split(w.Header().Get("isReturnRequested"), "")) > 0 {
		isReturnRequested, err := strconv.ParseBool(w.Header().Get("isReturnRequested"))
		Data.IsReturnRequested = isReturnRequested
		if err != nil {
			return err
		}
	}
	Data.UserId = int(userId)
	Data.IsAdmin = isAdmin
	Data.IsLoggedIn = IsLoggedIn
	return nil
}
