package convert

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

func blogFeedToVespaDocuments(blog BlogFeed) []VespaDocument {
	documents := make([]VespaDocument, len(blog.Posts))
	for i, post := range blog.Posts {
		fields := Fields{Language: post.getLanguage(), Id: post.getId(), Title: post.getTitle(), Body: post.getBody(), Url: post.getUrl(), Thumbnail: post.getThumbnail()}
		document := VespaDocument{Put: "id:myblog:myblog::" + post.getId(), Fields: fields}
		documents[i] = document
	}
	return documents
}

func XmlToJson(r io.Reader, w io.Writer) error {
	decoder := xml.NewDecoder(r)
	blogFeed := BlogFeed{}
	err := decoder.Decode(&blogFeed)
        if err != nil {
		return err
	}
	vespaDocuments := blogFeedToVespaDocuments(blogFeed)
	encoder := json.NewEncoder(w)
	encoder.SetEscapeHTML(false)
	err = encoder.Encode(vespaDocuments)
        if err != nil {
		return err
	}
	return nil
}
