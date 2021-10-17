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

func HttpHandleFunc(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	backendUrl := config.GetBackendUrl()
	proxy := httputil.NewSingleHostReverseProxy(backendUrl)
	rewriteRequest(req)
	log.Printf("Redirect %s%s?%s\n", backendUrl.String(), req.URL.Path, req.URL.RawQuery)
	proxy.ServeHTTP(res, req)
	
}
