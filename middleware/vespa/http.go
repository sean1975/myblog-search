package vespa

import (
	"github.com/sean1975/myblog-search/config"
	"log"
	"net/http"
	"strings"
)

func rewriteQueryStringForAutocomplete(req *http.Request) {
	query := req.URL.Query()
	_, ok := query["query"]
	if !ok {
		return
	}
	keyword := query.Get("query")
	query.Del("query")
	query.Set(
		"yql",
		"select title,url from sources * where title contains \""+keyword+"\";")
	req.URL.RawQuery = query.Encode()
}

func rewriteQueryStringForSearch(req *http.Request) {
	query := req.URL.Query()
	if _, ok := query["presentation.format"]; !ok {
		query.Add("presentation.format", "json")
	} else {
		query.Set("presentation.format", "json")
	}
	req.URL.RawQuery = query.Encode()
}

func rewriteQueryString(req *http.Request) {
	if strings.HasPrefix(req.URL.Path, "/autocomplete") {
		rewriteQueryStringForAutocomplete(req)
	} else {
		rewriteQueryStringForSearch(req)
	}
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

func RewriteRequest(req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	rewriteQueryString(req)
	rewriteRequestUrl(req)
	log.Printf("Redirect %s\n", req.URL.String())
}
