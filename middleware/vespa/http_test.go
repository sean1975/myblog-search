package vespa

import (
	"net/http"
	"strings"
	"testing"
)

func TestRewriteRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish", nil)
	rewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=json") != 1 {
		t.Errorf("Failed to rewrite request")
	}
}

func TestRewriteRequestXml(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish&presentation.format=xml", nil)
	rewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=json") != 1 {
		t.Errorf("Failed to rewrite request with XML format parameter")
	}
}

func TestRewriteRequestJson(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa/?query=fish&presentation.format=json", nil)
	rewriteRequest(req)
	if strings.Count(req.URL.String(), "presentation.format=json") != 1 {
		t.Errorf("Failed to rewrite request with JSON format parameter")
	}
}

func TestRewriteRequestUrl(t *testing.T) {
	req, _ := http.NewRequest("GET", "http://localhost/vespa", nil)
	rewriteRequestUrl(req)
	if req.URL.String() != "http://localhost/search/" {
		t.Errorf("Failed to rewrite request url " + req.URL.String())
	}
}

func TestNewResponseBody(t *testing.T) {
	body := `{"root":{"id":"toplevel","relevance":1.0,"fields":{"totalCount":1},"coverage":{"coverage":100,"documents":147,"full":true,"nodes":1,"results":1,"resultsFull":1},"children":[{"id":"id:myblog:myblog::2798327259924297827","relevance":0.08584490282742775,"source":"mycontent","fields":{"sddocname":"myblog","body":"<sep /> the first time divers have only an abstract idea about <hi>fish</hi> and coral. In general, there are at least three options<sep />","documentid":"id:myblog:myblog::2798327259924297827","id":"2798327259924297827","title":"Entertainment Options In An Introductory Dive","url":"http://blog.seanlee.site/2016/02/entertainment-options-in-introductory.html","thumbnail":"http://2.bp.blogspot.com/-EtQERDVbr50/Vq9STWInXeI/AAAAAAAAOsQ/_Vel08dEKGQ/s72-c/GOPR2092.JPG"}}]}}`
	expectedBody := `[{"id":"id:myblog:myblog::2798327259924297827","title":"Entertainment Options In An Introductory Dive","body":"... the first time divers have only an abstract idea about \u003cem\u003efish\u003c/em\u003e and coral. In general, there are at least three options...","url":"http://blog.seanlee.site/2016/02/entertainment-options-in-introductory.html","thumbnail":"http://2.bp.blogspot.com/-EtQERDVbr50/Vq9STWInXeI/AAAAAAAAOsQ/_Vel08dEKGQ/s72-c/GOPR2092.JPG"}]`
	buf := []byte(body)
	newBody := newResponseBody(buf)
	if string(newBody) != expectedBody {
		t.Errorf("Failed to create new response body: " + string(newBody))
	}
}
