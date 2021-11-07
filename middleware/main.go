package main

import (
	"github.com/sean1975/myblog-search/autocomplete"
	"github.com/sean1975/myblog-search/config"
	"github.com/sean1975/myblog-search/elastic"
	"github.com/sean1975/myblog-search/vespa"
	"net/http"
)

func main() {
	backendType := config.GetBackendType()
	vespaSearchHandler := vespa.NewSearchHandler()
	elasticSearchHandler := elastic.NewSearchHandler()
	autocompleteHandler := autocomplete.NewSearchHandler()
	if backendType == "vespa" {
		http.Handle("/search/", vespaSearchHandler)
	} else if backendType == "elastic" {
		http.Handle("/search/", elasticSearchHandler)
	}
	http.Handle("/vespa/", vespaSearchHandler)
	http.Handle("/elastic/", elasticSearchHandler)
	http.Handle("/autocomplete/", autocompleteHandler)
	if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}
