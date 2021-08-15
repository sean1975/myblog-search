package main

import (
    "net/http"
    "github.com/sean1975/myblog-search/config"
    "github.com/sean1975/myblog-search/search"
    "github.com/sean1975/myblog-search/autocomplete"
)

func main() {
    http.HandleFunc("/search/", search.HttpHandleFunc)
    http.HandleFunc("/autocomplete/", autocomplete.HttpHandleFunc)
    if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
        panic(err)
    }
}
