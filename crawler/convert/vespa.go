package convert

import (
	"encoding/json"
	"io"
)

type VespaDocument struct {
	Put string		`json:"put"`
	Fields VespaFields	`json:"fields"`
}

type VespaFields struct {
	Language string		`json:"language"`
	Id string		`json:"id"`
	Title string		`json:"title"`
	Body string		`json:"body"`
	Url string		`json:"url"`
	Thumbnail string	`json:"thumbnail"`
}

func newVespaDocuments(blog BlogFeed) []VespaDocument {
	documents := make([]VespaDocument, len(blog.Posts))
	for i, post := range blog.Posts {
		fields := VespaFields{Language: post.getLanguage(), Id: post.getId(), Title: post.getTitle(), Body: post.getBody(), Url: post.getUrl(), Thumbnail: post.getThumbnail()}
		document := VespaDocument{Put: "id:myblog:myblog::" + post.getId(), Fields: fields}
		documents[i] = document
	}
	return documents
}

type VespaEncoder struct {
	encoder			*json.Encoder
}

func NewVespaEncoder(w io.Writer) *VespaEncoder {
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetEscapeHTML(false)
	return &VespaEncoder{encoder: jsonEncoder}
}

func (enc *VespaEncoder) Encode(blog BlogFeed) error {
	documents := newVespaDocuments(blog)
	return enc.encoder.Encode(documents)
}
