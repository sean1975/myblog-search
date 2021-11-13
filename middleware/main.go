package main

import (
	"github.com/sean1975/myblog-search/config"
	elastic_autocomplete "github.com/sean1975/myblog-search/elastic/autocomplete"
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
	elasticAutocompleteHandler := elastic_autocomplete.NewHttpHandler()
	if backendType == "vespa" {
		http.Handle("/search/", vespaSearchHandler)
		http.Handle("/autocomplete/", vespaAutocompleteHandler)
	} else if backendType == "elastic" {
		http.Handle("/search/", elasticSearchHandler)
		http.Handle("/autocomplete/", elasticAutocompleteHandler)
	}
	http.Handle("/vespa/", vespaSearchHandler)
	http.Handle("/elastic/", elasticSearchHandler)
	if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}
