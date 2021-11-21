package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sean1975/myblog-search/vespa"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

type VespaSearchResult struct {
	Root Root `json:"root"`
}

type Root struct {
	Children []Child `json:"children"`
}

type Child struct {
	Id     string `json:"id"`
	Fields Fields `json:"fields"`
}

type Fields struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

type SearchResult struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

func newResponseBody(body []byte) []byte {
	var result VespaSearchResult
	json.Unmarshal(body, &result)
	var newResult = make([]SearchResult, len(result.Root.Children))
	for i, child := range result.Root.Children {
		newResult[i].Id = child.Id
		newResult[i].Title = child.Fields.Title
		newResult[i].Url = child.Fields.Url
		newResult[i].Thumbnail = child.Fields.Thumbnail
		newResult[i].Body = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(child.Fields.Body, "\u003csep /\u003e", "..."), "\u003chi\u003e", "\u003cem\u003e"), "\u003c/hi\u003e", "\u003c/em\u003e")
	}
	newBody, _ := json.Marshal(newResult)
	return newBody
}

func rewriteResponseBody(res *http.Response) {
	body, _ := ioutil.ReadAll(res.Body)
	newBody := newResponseBody(body)
	buf := bytes.NewBuffer(newBody)
	res.Body = ioutil.NopCloser(buf)
	res.Header["Content-Length"] = []string{fmt.Sprint(buf.Len())}
}

func rewriteResponse(res *http.Response) error {
	if res.StatusCode == 200 {
		rewriteResponseBody(res)
	}
	return nil
}

func NewHttpHandler() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: vespa.RewriteRequest, ModifyResponse: rewriteResponse}
}
