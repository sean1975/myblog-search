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
	"net/url"
	"strings"
)

func rewriteQueryValues(query url.Values) url.Values {
	_, ok := query["query"]
	if !ok {
		return query
	}
	keyword := query.Get("query")
	query.Del("query")
	query.Set(
		"yql",
		"select title,url from sources * where title contains \""+keyword+"\";")
	return query
}

func rewriteRequestUrl(req *http.Request) {
	backendUrl := config.GetBackendUrl()
	if !strings.HasSuffix(backendUrl.EscapedPath(), "/") {
		backendUrl.Path += "/"
	}
	backendUrl.Path += "search/"
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
	query := req.URL.Query()
	req.URL.RawQuery = rewriteQueryValues(query).Encode()
	rewriteRequestUrl(req)
	log.Printf("Redirect %s\n", req.URL.String())
}

type Result struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func rewriteResponseBody(responseBody []byte) ([]byte, error) {
	results := make([]Result, 0)
	var result map[string]interface{}
	err := json.Unmarshal(responseBody, &result)
	if err != nil {
		return json.Marshal(results)
	}
	root := result["root"].(map[string]interface{})
	if children, ok := root["children"].([]interface{}); ok {
		for i := range children {
			child := children[i].(map[string]interface{})
			fields := child["fields"].(map[string]interface{})
			title := fields["title"].(string)
			url := fields["url"].(string)
			results = append(results, Result{title, url})
		}
	}
	return json.Marshal(results)
}

func rewriteResponse(res *http.Response) error {
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		newBody, _ := rewriteResponseBody(body)
		buf := bytes.NewBuffer(newBody)
		res.Body = ioutil.NopCloser(buf)
		res.Header["Content-Length"] = []string{fmt.Sprint(buf.Len())}
	}
	return nil
}

func NewHttpHandler() *httputil.ReverseProxy {
	return &httputil.ReverseProxy{Director: rewriteRequest, ModifyResponse: rewriteResponse}
}
