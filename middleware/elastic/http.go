package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sean1975/myblog-search/config"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
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

type SearchTemplate struct {
	Id string		`json:"id"`
	Params Params		`json:"params"`
}

type Params struct {
	QueryString string	`json:"query_string"`
}

func newRequestBody(queryString string) []byte {
	param := Params{QueryString: queryString}
	bodyObject := SearchTemplate{Id: "myblog-search-template", Params: param}
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
	req.URL.Path = "/_search/template"
}

func rewriteRequest(req *http.Request) {
	rewriteRequestBody(req)
	removeQueryString(req)
	rewriteRequestUrl(req)
}

type ElasticSearchResult struct {
	Hits Hits			`json:"hits"`
}

type Hits struct {
	InnerHits []InnerHit		`json:"hits"`
}

type InnerHit struct {
	Index string			`json:"_index"`
	Type string			`json:"_type"`
	Id string			`json:"_id"`
	Score float64			`json:"_score"`
	Fields Fields			`json:"fields"`
	Highlight Highlight		`json:"highlight"`
}

type Fields struct {
	Title []string			`json:"title"`
	Url []string			`json:"url"`
	Thumbnail []string		`json:"thumbnail"`
}

type Highlight struct {
	Body []string			`json:"body"`
}

type SearchResult struct {
	Id string			`json:"id"`
	Title string			`json:"title"`
	Body string			`json:"body"`
	Url string			`json:"url"`
	Thumbnail string		`json:"thumbnail"`
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

func HttpHandleFunc(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	backendUrl := config.GetBackendUrl()
	proxy := httputil.NewSingleHostReverseProxy(backendUrl)
	proxy.ModifyResponse = rewriteResponse
	rewriteRequest(req)
	log.Printf("Redirect %s%s?%s\n", backendUrl.String(), req.URL.Path, req.URL.RawQuery)
	proxy.ServeHTTP(res, req)
	
}
