package utils

import (
	"net/url"
	"strings"
)

func GetPageNumberFromQueryString(queryString url.Values) string {
	page := strings.TrimSpace(queryString.Get("page"))
	if page == "" {
		return "1"
	}
	return page
}
