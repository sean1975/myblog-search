package vespa

import (
	"github.com/sean1975/myblog-search/config"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func appendQueryValues(query url.Values) url.Values {
	if _, ok := query["presentation.format"]; !ok {
		query.Add("presentation.format", "xml")
	}
	return query
}

func rewriteRequestUrl(req *http.Request) {
	req.URL.Path = "/search"
}

func rewriteRequest(req *http.Request) {
	query := req.URL.Query()
	req.URL.RawQuery = appendQueryValues(query).Encode()
}

func HttpHandleFunc(res http.ResponseWriter, req *http.Request) {
	log.Printf("Request %s\n", req.URL.String())
	backendUrl := config.GetBackendUrl()
	proxy := httputil.NewSingleHostReverseProxy(backendUrl)
	rewriteRequest(req)
	log.Printf("Redirect %s%s?%s\n", backendUrl.String(), req.URL.Path, req.URL.RawQuery)
	proxy.ServeHTTP(res, req)
}
