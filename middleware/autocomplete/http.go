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

func rewriteRequest(req *http.Request) {
	query := req.URL.Query()
	req.URL.RawQuery = rewriteQueryValues(query).Encode()
	req.URL.Path = strings.Replace(req.URL.Path, "autocomplete", "search", 1)
}

type Result struct {
	Title string `json:title`
	Url   string `json:url`
}

func rewriteResponseBody(responseBody []byte) ([]byte, error) {
	var result map[string]interface{}
	json.Unmarshal(responseBody, &result)
	root := result["root"].(map[string]interface{})
	results := make([]Result, 0)
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

func HttpHandleFunc(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	backendUrl := config.GetBackendUrl()
	proxy := httputil.NewSingleHostReverseProxy(backendUrl)
	proxy.ModifyResponse = rewriteResponse
	rewriteRequest(req)
	log.Printf("Redirect %s%s?%s\n", backendUrl.String(), req.URL.Path, req.URL.RawQuery)
	proxy.ServeHTTP(res, req)
}
