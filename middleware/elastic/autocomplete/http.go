package autocomplete

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sean1975/myblog-search/config"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

func getQueryString(req *http.Request) string {
	query := req.URL.Query()
	_, ok := query["query"]
	if !ok {
		return ""
	}
	return query.Get("query")
}

func removeQueryString(req *http.Request) {
	query := req.URL.Query()
	_, ok := query["query"]
	if ok {
		query.Del("query")
	}
	req.URL.RawQuery = query.Encode()
	req.URL.ForceQuery = false
}

type AutocompleteTemplate struct {
	Id     string `json:"id"`
	Params Params `json:"params"`
}

type Params struct {
	QueryString string `json:"query_string"`
}

func newRequestBody(queryString string) []byte {
	param := Params{QueryString: queryString}
	bodyObject := AutocompleteTemplate{Id: "myblog-autocomplete-template", Params: param}
	body, _ := json.Marshal(bodyObject)
	return body
}

func rewriteRequestBody(req *http.Request) {
	queryString := getQueryString(req)
	body := newRequestBody(queryString)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	req.ContentLength = int64(len(body))
	req.Header["Content-Length"] = []string{fmt.Sprint(len(body))}
	req.Header["Content-Type"] = []string{"application/json"}
}

func rewriteRequestUrl(req *http.Request) {
	backendUrl := config.GetBackendUrl()
	if !strings.HasSuffix(backendUrl.EscapedPath(), "/") {
		backendUrl.Path += "/"
	}
	backendUrl.Path += "_search/template"
	req.URL.Scheme = backendUrl.Scheme
	req.URL.Host = backendUrl.Host
	req.URL.Path = backendUrl.Path
	req.URL.RawPath = backendUrl.Path
	if _, ok := req.Header["User-Agent"]; !ok {
		// explicitly disable User-Agent so it's not set to default value
		req.Header.Set("User-Agent", "")
	}
}

func rewriteRequest(req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	rewriteRequestBody(req)
	removeQueryString(req)
	rewriteRequestUrl(req)
	log.Printf("Redirect %s\n", req.URL.String())
}

type ElasticSearchResult struct {
	Hits Hits `json:"hits"`
}

type Hits struct {
	InnerHits []InnerHit `json:"hits"`
}

type InnerHit struct {
	Index  string  `json:"_index"`
	Type   string  `json:"_type"`
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Fields Fields  `json:"fields"`
}

type Fields struct {
	Title []string `json:"title"`
	Url   []string `json:"url"`
}

type SearchResult struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func newResponseBody(body []byte) []byte {
	var result ElasticSearchResult
	json.Unmarshal(body, &result)
	var newResult = make([]SearchResult, len(result.Hits.InnerHits))
	for i, hit := range result.Hits.InnerHits {
		newResult[i].Title = hit.Fields.Title[0]
		newResult[i].Url = hit.Fields.Url[0]
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
	return &httputil.ReverseProxy{Director: rewriteRequest, ModifyResponse: rewriteResponse}
}
