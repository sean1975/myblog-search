package vespa

import (
	"net/http"
	"strings"
	"testing"
)

func TestRewriteRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish", nil)
	RewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=json") != 1 {
		t.Errorf("Failed to rewrite request")
	}
}

func TestRewriteRequestXml(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish&presentation.format=xml", nil)
	RewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=json") != 1 {
		t.Errorf("Failed to rewrite request with XML format parameter")
	}
}

func TestRewriteRequestJson(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish&presentation.format=json", nil)
	RewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=json") != 1 {
		t.Errorf("Failed to rewrite request with JSON format parameter")
	}
}

func TestRewriteRequestUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa", nil)
	rewriteRequestUrl(req)
	if req.URL.String() != "http://localhost:8080/search/" {
		t.Errorf("Failed to rewrite request url " + req.URL.String())
	}
}

func TestRewriteRequestAutocomplete(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/autocomplete/?query=fish", nil)
	RewriteRequest(req)
	if !strings.Contains(req.URL.Path, "/search/") {
		t.Errorf("Failed to rewrite request - incorrect path %s", req.URL.Path)
	}
	url := req.URL.String()
	if !strings.Contains(url, "?") {
		t.Errorf("Failed to rewrite request - no query parameter")
	}
	queryParameters := strings.SplitN(url, "?", 2)
	if !strings.HasPrefix(queryParameters[1], "yql") {
		t.Errorf("Failed to rewrite request - incorrect query parameter %s",
			queryParameters[1])
	}
}
