package main

import (
	"github.com/sean1975/myblog-search/config"
	elastic_search "github.com/sean1975/myblog-search/elastic/search"
	vespa_autocomplete "github.com/sean1975/myblog-search/vespa/autocomplete"
	vespa_search "github.com/sean1975/myblog-search/vespa/search"
	"net/http"
)

func main() {
	backendType := config.GetBackendType()
	vespaSearchHandler := vespa_search.NewHttpHandler()
	elasticSearchHandler := elastic_search.NewHttpHandler()
	vespaAutocompleteHandler := vespa_autocomplete.NewHttpHandler()
	if backendType == "vespa" {
		http.Handle("/search/", vespaSearchHandler)
	} else if backendType == "elastic" {
		http.Handle("/search/", elasticSearchHandler)
	}
	http.Handle("/vespa/", vespaSearchHandler)
	http.Handle("/elastic/", elasticSearchHandler)
	http.Handle("/autocomplete/", vespaAutocompleteHandler)
	if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}
