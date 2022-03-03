package web

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

// mergePathOrQueryString from original request for more flexible usage.
// For example:
// Assume we have record: go/issue -> http://ticket.com
// this function extends the redirection with following ops
// 1. go/issue/ticket-001 -> http://ticket.com/ticket-001
// 2. go/issue?ticket-001 -> http://ticket.com?ticket-001
//
// Since this is an optional flow, the returned value is implicit set to rtURL if
// error happends.
func mergePathOrQueryString(p, rtURL string, r *http.Request) string {
	noSubPath := strings.HasSuffix(strings.TrimSuffix(r.URL.Path, "/"), p)
	noQueryString := len(r.URL.Query()) == 0

	if noSubPath && noQueryString {
		return rtURL
	}

	newURL, err := url.Parse(rtURL)
	if err != nil {
		return rtURL
	}

	newURL.RawQuery = r.URL.RawQuery
	subPath := parseSubPath("/", r.URL.Path)
	newURL.Path = filepath.Join(newURL.Path, subPath)

	return newURL.String()
}
