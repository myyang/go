package web

import (
	"net/http"
	"net/url"
	"testing"
)

func TestMergePathOrQueryString(t *testing.T) {
	request := &http.Request{}
	p, rtURL := "short", "http://target.com"

	tests := []struct {
		Case       string
		RequestURL string
		Exp        string
	}{
		{"normal", "http://go/short/", "http://target.com"},
		{"normal - more tailing slashs", "http://go/short/////", "http://target.com/"}, // forced root path
		{"normal - no tailing slash", "http://go/short", "http://target.com"},
		{"with sub path", "http://go/short/abc", "http://target.com/abc"},
		{"with sub paths - i", "http://go/short/abc/def", "http://target.com/abc/def"},
		{"with sub paths - ii", "http://go/short/////abc/def", "http://target.com/abc/def"},
		{"with query string", "http://go/short?a=b", "http://target.com?a=b"},
		{"with query strings", "http://go/short?a=b&c=d", "http://target.com?a=b&c=d"},
		{"with sub paths and query strings", "http://go/short/abc/def?a=b&c=d", "http://target.com/abc/def?a=b&c=d"},
	}

	for _, test := range tests {
		tURL, err := url.Parse(test.RequestURL)
		if err != nil {
			t.Fatalf("case: %v. parse rquest URL error: %v", test.Case, err)
		}

		request.URL = tURL

		exp := mergePathOrQueryString(p, rtURL, request)
		if test.Exp != exp {
			t.Fatalf("case %v. not expected result: %v != %v", test.Case, test.Exp, exp)
		}
	}
}
