package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sean1975/myblog-search/config"
	"io/ioutil"
	"log"
	"net/http"
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

type ElasticSearchTemplate struct {
	Id     string `json:"id"`
	Params Params `json:"params"`
}

type Params struct {
	QueryString string `json:"query_string"`
}

func newRequestBody(templateId, queryString string) []byte {
	param := Params{QueryString: queryString}
	bodyObject := ElasticSearchTemplate{Id: templateId, Params: param}
	body, _ := json.Marshal(bodyObject)
	return body
}

func getTemplateIdFromRequestPath(req *http.Request) string {
	templateId := "myblog-search-template"
	if strings.HasPrefix(req.URL.Path, "/autocomplete") {
		templateId = "myblog-autocomplete-template"
	}
	return templateId
}

func rewriteRequestBody(req *http.Request) {
	queryString := getQueryString(req)
	body := newRequestBody(getTemplateIdFromRequestPath(req), queryString)
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

func RewriteRequest(req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	rewriteRequestBody(req)
	removeQueryString(req)
	rewriteRequestUrl(req)
	log.Printf("Redirect %s\n", req.URL.String())
}
