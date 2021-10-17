package main

import (
	"github.com/sean1975/myblog-search/autocomplete"
	"github.com/sean1975/myblog-search/config"
	"github.com/sean1975/myblog-search/vespa"
	"github.com/sean1975/myblog-search/elastic"
	"net/http"
)

func main() {
	backendType := config.GetBackendType();
	if (backendType == "vespa") {
		http.HandleFunc("/search/", vespa.HttpHandleFunc)
	} else if (backendType == "elastic") {
		http.HandleFunc("/search/", elastic.HttpHandleFunc)
	}
	http.HandleFunc("/vespa/", vespa.HttpHandleFunc)
	http.HandleFunc("/elastic/", elastic.HttpHandleFunc)
	http.HandleFunc("/autocomplete/", autocomplete.HttpHandleFunc)
	if err := http.ListenAndServe(config.GetListenAddress(), nil); err != nil {
		panic(err)
	}
}
