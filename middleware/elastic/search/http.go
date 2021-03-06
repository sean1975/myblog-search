package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sean1975/myblog-search/elastic"
	"github.com/sean1975/myblog-search/utils"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

type ElasticSearchResult struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	InnerHits []InnerHit `json:"hits"`
}

type InnerHit struct {
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	Id        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Fields    Fields    `json:"fields"`
	Highlight Highlight `json:"highlight"`
}

type Fields struct {
	Title     []string `json:"title"`
	Url       []string `json:"url"`
	Thumbnail []string `json:"thumbnail"`
}

type Highlight struct {
	Body []string `json:"body"`
}

type SearchResult struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Url       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
}

func newResponseBody(body []byte) []byte {
	var result ElasticSearchResult
	json.Unmarshal(body, &result)
	var newResult = make([]SearchResult, len(result.Hits.InnerHits))
	for i, hit := range result.Hits.InnerHits {
		newResult[i].Id = hit.Id
		newResult[i].Title = hit.Fields.Title[0]
		newResult[i].Url = hit.Fields.Url[0]
		newResult[i].Thumbnail = hit.Fields.Thumbnail[0]
		newResult[i].Body = hit.Highlight.Body[0]
	}
	newBody, _ := json.Marshal(newResult)
	return newBody
}

func rewriteResponseBody(res *http.Response) {
	body, _ := ioutil.ReadAll(res.Body)
	if res.Header.Get("Content-Encoding") == "gzip" {
		body = utils.Gunzip(body)
	}
	body = newResponseBody(body)
	if res.Header.Get("Content-Encoding") == "gzip" {
		body = utils.Gzip(body)
	}
	buf := bytes.NewBuffer(body)
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
	return &httputil.ReverseProxy{Director: elastic.RewriteRequest, ModifyResponse: rewriteResponse}
}
