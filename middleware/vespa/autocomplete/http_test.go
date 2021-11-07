package autocomplete

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestRewriteRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/autocomplete/?query=fish", nil)
	rewriteRequest(req)
	if !strings.Contains(req.URL.Path, "/search/") {
		t.Errorf("Failed to rewrite request - incorrect path %s", req.URL.Path)
	}
	url := req.URL.String()
	if !strings.Contains(url, "?") {
		t.Errorf("Failed to rewrite request - no query parameter")
	}
	queryParameters := strings.SplitN(url, "?", 2)
	if !strings.HasPrefix(queryParameters[1], "yql") {
		t.Errorf("Failed to rewrite request - incorrect query parameter %s",
			queryParameters[1])
	}
}

func TestRewriteResponse(t *testing.T) {
	body := `{"root":{"id":"toplevel","relevance":1.0,"fields":{"totalCount":1},"coverage":{"coverage":100,"documents":145,"full":true,"nodes":1,"results":1,"resultsFull":1},"children":[{"id":"index:mycontent/0/91ba7c3a0633aeabd9878503","relevance":0.16343879032006287,"source":"mycontent","fields":{"title":"很多魚","url":"http://blog.seanlee.site/2014/08/blog-post_16.html"}}]}}`
	res := &http.Response{
		StatusCode: 200,
		Header:     http.Header{},
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}
	rewriteResponse(res)
	newBody, _ := ioutil.ReadAll(res.Body)
	var results []Result
	json.Unmarshal(newBody, &results)
	if len(results) != 1 {
		t.Errorf("Failed to rewrite response - incorrect result number %d", len(results))
	}
	result := results[0]
	if result.Title != "很多魚" {
		t.Errorf("Failed to rewrite response - incorrect title %s", result.Title)
	}
	if result.Url != "http://blog.seanlee.site/2014/08/blog-post_16.html" {
		t.Errorf("Failed to rewrite response - incorrect url %s", result.Url)
	}
}

func TestRewriteResponseEmpty(t *testing.T) {
	body := `{"root":{"id":"toplevel","relevance":1.0,"fields":{"totalCount":0},"coverage":{"coverage":100,"documents":145,"full":true,"nodes":1,"results":1,"resultsFull":1}}}`
	res := &http.Response{
		StatusCode: 200,
		Header:     http.Header{},
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}
	rewriteResponse(res)
	newBody, _ := ioutil.ReadAll(res.Body)
	var results []Result
	json.Unmarshal(newBody, &results)
	if len(results) != 0 {
		t.Errorf("Failed to rewrite response - incorrect result number %d", len(results))
	}
}
