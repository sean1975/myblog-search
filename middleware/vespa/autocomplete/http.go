package autocomplete

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sean1975/myblog-search/vespa"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

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
	return &httputil.ReverseProxy{Director: vespa.RewriteRequest, ModifyResponse: rewriteResponse}
}
