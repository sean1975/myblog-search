package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "os"
)

func appendQueryValues(query url.Values) url.Values {
    if _, ok := query["presentation.format"]; !ok {
        query.Add("presentation.format", "xml")
    }
    return query
}

func rewriteRequest(req *http.Request) {
    query := req.URL.Query()
    req.URL.RawQuery = appendQueryValues(query).Encode()
}

func searchHandler(res http.ResponseWriter, req *http.Request) {
    log.Printf("Request %s\n", req.URL.String())
    backendUrl := getBackendUrl()
    proxy := httputil.NewSingleHostReverseProxy(backendUrl)
    rewriteRequest(req)
    log.Printf("Redirect %s%s?%s\n", backendUrl.String(), req.URL.Path, req.URL.RawQuery)
    proxy.ServeHTTP(res, req)
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func getBackendUrl() *url.URL {
    backendEnv := getEnv("BACKEND_URL", "http://localhost:8080")
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
