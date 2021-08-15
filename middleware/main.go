package main

import (
	"github.com/sean1975/myblog-search/autocomplete"
	"github.com/sean1975/myblog-search/config"
	"github.com/sean1975/myblog-search/search"
	"net/http"
)

func main() {
	http.HandleFunc("/search/", search.HttpHandleFunc)
	http.HandleFunc("/autocomplete/", autocomplete.HttpHandleFunc)
	if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}
