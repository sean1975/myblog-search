package elastic 

import (
	"net/http"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRemoveQueryString(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/elastic/?query=fish", nil)
	removeQueryString(req)
	if strings.Count(req.URL.String(), "query=fish") != 0 {
		t.Errorf("Failed to remove query string: " + req.URL.String())
	}
}

func TestRewriteRequestBody(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/elastic/?query=fish", nil)
	expectedBody := `{"id":"myblog-search-template","params":{"query_string":"fish"}}`
	rewriteRequestBody(req)
	if req.ContentLength != int64(len(expectedBody)) {
		t.Errorf("Content length is incorrect after rewrite request body")
	}
	newBody, _ := ioutil.ReadAll(req.Body)
	if string(newBody) != expectedBody {
		t.Errorf("Failed to rewrite request body: " + string(newBody))
	}
}

func TestRewriteRequestUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/elastic", nil)
	rewriteRequestUrl(req)
	if req.URL.String() != "http://localhost/_search/template" {
		t.Errorf("Failed to rewrite request url: " + req.URL.String())
	}
}
