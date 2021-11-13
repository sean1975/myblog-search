package autocomplete

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRemoveQueryString(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/autocomplete/?query=fish", nil)
	removeQueryString(req)
	if strings.Count(req.URL.String(), "query=fish") != 0 {
		t.Errorf("Failed to remove query string: " + req.URL.String())
	}
}

func TestRewriteRequestBody(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/autocomplete/?query=div", nil)
	expectedBody := `{"id":"myblog-autocomplete-template","params":{"query_string":"div"}}`
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
	if req.URL.String() != "http://localhost:8080/_search/template" {
		t.Errorf("Failed to rewrite request url: " + req.URL.String())
	}
}

func TestNewResponseBody(t *testing.T) {
	body := `{"took":19,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":6.3361177,"hits":[{"_index":"myblog","_type":"_doc","_id":"5367749013774290929","_score":6.3361177,"fields":{"title":["Panic Diver"],"url":["http://blog.seanlee.site/2016/03/panic-diver.html"]}}]}}`
	expectedBody := `[{"title":"Panic Diver","url":"http://blog.seanlee.site/2016/03/panic-diver.html"}]`
	buf := []byte(body)
	newBody := newResponseBody(buf)
	if string(newBody) != expectedBody {
		t.Errorf("Failed to create new response body: " + string(newBody))
	}
}
