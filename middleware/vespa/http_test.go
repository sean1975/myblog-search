package vespa

import (
	"net/http"
	"strings"
	"testing"
)

func TestRewriteRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish", nil)
	rewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=xml") != 1 {
		t.Errorf("Failed to rewrite request")
	}
}

func TestRewriteRequestXml(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish&presentation.format=xml", nil)
	rewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=xml") != 1 {
		t.Errorf("Failed to rewrite request with XML format parameter")
	}
}

func TestRewriteRequestJson(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish&presentation.format=json", nil)
	rewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=xml") != 0 {
		t.Errorf("Failed to rewrite request with JSON format parameter")
	}
}

func TestRewriteRequestUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa", nil)
	rewriteRequestUrl(req)
	if req.URL.String() != "http://localhost/search" {
		t.Errorf("Failed to rewrite request url " + req.URL.String())
	}
}

