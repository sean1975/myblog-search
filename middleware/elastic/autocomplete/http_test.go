package autocomplete

import (
	"testing"
)

func TestNewResponseBody(t *testing.T) {
	body := `{"took":19,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":6.3361177,"hits":[{"_index":"myblog","_type":"_doc","_id":"5367749013774290929","_score":6.3361177,"fields":{"title":["Panic Diver"],"url":["http://blog.seanlee.site/2016/03/panic-diver.html"]}}]}}`
	expectedBody := `[{"title":"Panic Diver","url":"http://blog.seanlee.site/2016/03/panic-diver.html"}]`
	buf := []byte(body)
	newBody := newResponseBody(buf)
	if string(newBody) != expectedBody {
		t.Errorf("Failed to create new response body: " + string(newBody))
	}
}
