package convert

import (
	"bytes"
	"strings"
	"testing"
)

func TestXmlToElasticJson(t *testing.T) {
	xmlData := `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="http://www.blogger.com/styles/atom.css" type="text/css"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:openSearch="http://a9.com/-/spec/opensearchrss/1.0/" xmlns:blogger="http://schemas.google.com/blogger/2008" xmlns:georss="http://www.georss.org/georss" xmlns:gd="http://schemas.google.com/g/2005" xmlns:thr="http://purl.org/syndication/thread/1.0">
  <updated>2021-08-24T16:31:24.616+08:00</updated>
  <entry>
    <id>tag:blogger.com,1999:blog-8781445187228811595.post-3214139340207317638</id>
    <published>2020-10-29T18:26:00.011+08:00</published>
    <updated>2020-12-02T20:20:57.441+08:00</updated>
    <category scheme="http://www.blogger.com/atom/ns#" term="Australia"/>
    <category scheme="http://www.blogger.com/atom/ns#" term="Taiwan"/>
    <title type="text">Blog Post Title</title>
    <content type="html">&amp;nbsp; &amp;nbsp; Sentence1.&lt;br /&gt;&amp;nbsp; &amp;nbsp; Sentence2.&lt;br /&gt;&amp;nbsp; &amp;nbsp; Sentence3.&lt;br /&gt;&lt;div class="separator" style="clear: both;"&gt;&lt;a href="https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s0/20201123_125051.jpg" style="display: block; padding: 1em 0; text-align: center; "&gt;&lt;img alt="" border="0" with="640" data-original-height="1536" data-original-width="2048" src="https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s0/20201123_125051.jpg"/&gt;&lt;/a&gt;&lt;/div&gt;</content>
    <link rel="self" type="application/atom+xml" href="http://www.blogger.com/feeds/8781445187228811595/posts/default/3214139340207317638"/>
    <link rel="alternate" type="text/html" href="http://blog.seanlee.site/2020/10/congratulations-your-application-has.html" title="Congratulations, your application has been approved"/>
    <media:thumbnail xmlns:media="http://search.yahoo.com/mrss/" url="https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s72-c/20201123_125051.jpg" height="72" width="72"/>
  </entry>
</feed>
`
	jsonData := `{"index":{"_index":"myblog","_id":"3214139340207317638"}}
{"language":"en","title":"Blog Post Title","body":"Sentence1. Sentence2. Sentence3.","url":"http://blog.seanlee.site/2020/10/congratulations-your-application-has.html","thumbnail":"https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s72-c/20201123_125051.jpg"}`
	r := strings.NewReader(xmlData)
	buf := new(bytes.Buffer)
	err := XmlToJson(r, buf, "elastic")
	if err != nil {
		t.Errorf("Failed to convert XML into Json, error %s", err)
	}
	if strings.TrimSpace(buf.String()) != strings.TrimSpace(jsonData) {
		t.Errorf("Incorrect Json data:\n%s\nExpected Json data:\n%s\n", buf.String(), jsonData)
	}
}

func TestXmlToElasticJsonChinese(t *testing.T) {
	xmlData := `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet href="http://www.blogger.com/styles/atom.css" type="text/css"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:openSearch="http://a9.com/-/spec/opensearchrss/1.0/" xmlns:blogger="http://schemas.google.com/blogger/2008" xmlns:georss="http://www.georss.org/georss" xmlns:gd="http://schemas.google.com/g/2005" xmlns:thr="http://purl.org/syndication/thread/1.0">
  <updated>2021-08-24T16:31:24.616+08:00</updated>
  <entry>
    <id>tag:blogger.com,1999:blog-8781445187228811595.post-3214139340207317638</id>
    <published>2020-10-29T18:26:00.011+08:00</published>
    <updated>2020-12-02T20:20:57.441+08:00</updated>
    <category scheme="http://www.blogger.com/atom/ns#" term="Taiwan"/>
    <title type="text">Blog Post Title</title>
    <content type="html">&lt;div style="text-align: right;"&gt;&lt;a href="http://diaryofsean.blogspot.com/2020/10/congratulations-your-application-has.html"&gt;English version&lt;/a&gt;&lt;/div&gt;&amp;nbsp; &amp;nbsp; Sentence1.&lt;br /&gt;&amp;nbsp; &amp;nbsp; Sentence2.&lt;br /&gt;&amp;nbsp; &amp;nbsp; Sentence3.&lt;br /&gt;&lt;div class="separator" style="clear: both;"&gt;&lt;a href="https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s0/20201123_125051.jpg" style="display: block; padding: 1em 0; text-align: center; "&gt;&lt;img alt="" border="0" with="640" data-original-height="1536" data-original-width="2048" src="https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s0/20201123_125051.jpg"/&gt;&lt;/a&gt;&lt;/div&gt;</content>
    <link rel="self" type="application/atom+xml" href="http://www.blogger.com/feeds/8781445187228811595/posts/default/3214139340207317638"/>
    <link rel="alternate" type="text/html" href="http://blog.seanlee.site/2020/10/congratulations-your-application-has.html" title="Congratulations, your application has been approved"/>
    <media:thumbnail xmlns:media="http://search.yahoo.com/mrss/" url="https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s72-c/20201123_125051.jpg" height="72" width="72"/>
  </entry>
</feed>
`
	jsonData := `{"index":{"_index":"myblog","_id":"3214139340207317638"}}
{"language":"zh-TW","title":"Blog Post Title","body":"Sentence1. Sentence2. Sentence3.","url":"http://blog.seanlee.site/2020/10/congratulations-your-application-has.html","thumbnail":"https://1.bp.blogspot.com/-VbZs62V8a9U/X8eDJBbbBWI/AAAAAAAAja4/jj8kbf5p-6YvDXx8FDoIJ27NB0K4-YZFACLcBGAsYHQ/s72-c/20201123_125051.jpg"}`
	r := strings.NewReader(xmlData)
	buf := new(bytes.Buffer)
	err := XmlToJson(r, buf, "elastic")
	if err != nil {
		t.Errorf("Failed to convert XML into Json, error %s", err)
	}
	if strings.TrimSpace(buf.String()) != strings.TrimSpace(jsonData) {
		t.Errorf("Incorrect Json data:\n%s\nExpected Json data:\n%s\n", buf.String(), jsonData)
	}
}
