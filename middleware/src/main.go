package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
    "strings"
)

func filterEmptyString(vals []string) []string {
    n := 0
    for _, val := range vals {
        if val != "" {
            vals[n] = val
            n++
        }
    }
    vals = vals[:n]
    return vals
}

func appendQueryValues(query url.Values, vals []string) url.Values {
    for _, val := range vals {
        query.Add("query", val)
    }
    if _, ok := query["presentation.format"]; !ok {
        query.Add("presentation.format", "xml")
    }
    return query
}

func rewriteRequestQuery(req *http.Request) {
    keywords := strings.Split(req.URL.Path, "/")
    keywords = filterEmptyString(keywords)
    query := req.URL.Query()
    req.URL.RawQuery = appendQueryValues(query, keywords).Encode()
}

func rewriteRequest(req *http.Request) {
    rewriteRequestQuery(req)
    req.URL.Path = "/"
}

func searchHandler(res http.ResponseWriter, req *http.Request) {
    log.Printf("Request %s\n", req.URL.String())
    backendUrl := getBackendUrl()
    proxy := httputil.NewSingleHostReverseProxy(backendUrl)
    rewriteRequest(req)
    log.Printf("Redirect %s?%s\n", backendUrl.String(), req.URL.RawQuery)
    proxy.ServeHTTP(res, req)
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func getBackendUrl() *url.URL {
    backendEnv := getEnv("BACKEND_URL", "http://localhost:8080/search")
    backendUrl, err := url.Parse(backendEnv)
    if err != nil {
        panic(err)
    }
    return backendUrl
}
    
func getListenAddress() string {
    port := getEnv("PORT", "80")
    return ":" + port
}

func main() {
    http.HandleFunc("/", searchHandler)
    if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
        panic(err)
    }
}
