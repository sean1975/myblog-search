package search

import (
	"testing"
)

func TestNewResponseBody(t *testing.T) {
	body := `{"took":105,"timed_out":false,"_shards":{"total":1,"successful":1,"skipped":0,"failed":0},"hits":{"total":{"value":1,"relation":"eq"},"max_score":5.7349234,"hits":[{"_index":"myblog","_type":"_doc","_id":"2798327259924297827","_score":5.7349234,"fields":{"thumbnail":["http://2.bp.blogspot.com/-EtQERDVbr50/Vq9STWInXeI/AAAAAAAAOsQ/_Vel08dEKGQ/s72-c/GOPR2092.JPG"],"title":["Entertainment Options In An Introductory Dive"],"url":["http://blog.seanlee.site/2016/02/entertainment-options-in-introductory.html"]},"highlight":{"body":["Especially the first time divers have only an abstract idea about <em>fish</em> and coral.","They used to be called anemone <em>fish</em> or clown <em>fish</em>.","Due to the film Finding Nemo, the <em>fish</em> changed its name overnight in all languages.","the Nemo at a prefect location where customers can sit on the seabed for closely watching the small <em>fish</em>"]}}]}}`
	expectedBody := `[{"id":"2798327259924297827","title":"Entertainment Options In An Introductory Dive","body":"Especially the first time divers have only an abstract idea about \u003cem\u003efish\u003c/em\u003e and coral.","url":"http://blog.seanlee.site/2016/02/entertainment-options-in-introductory.html","thumbnail":"http://2.bp.blogspot.com/-EtQERDVbr50/Vq9STWInXeI/AAAAAAAAOsQ/_Vel08dEKGQ/s72-c/GOPR2092.JPG"}]`
	buf := []byte(body)
	newBody := newResponseBody(buf)
	if string(newBody) != expectedBody {
		t.Errorf("Failed to create new response body: " + string(newBody))
	}
}
