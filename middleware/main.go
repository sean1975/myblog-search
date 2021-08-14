package main

import (
    "net/http"
    "github.com/sean1975/myblog-search/config"
    "github.com/sean1975/myblog-search/search"
)

func main() {
    http.HandleFunc("/", search.HttpHandleFunc)
    if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
        panic(err)
    }
}
