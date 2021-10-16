package convert

import (
	"encoding/json"
	"io"
)

type ElasticDocument struct {
	Action ElasticAction
	Fields ElasticFields
}

type ElasticAction struct {
	Index IndexMeta		`json:"index"`
}

type IndexMeta struct {
	IndexName string	`json:"_index"`
	Id string		`json:"_id"`
}

type ElasticFields struct {
	Language string		`json:"language"`
	Title string		`json:"title"`
	Body string		`json:"body"`
	Url string		`json:"url"`
	Thumbnail string	`json:"thumbnail"`
}

func newElasticDocuments(blog BlogFeed) []ElasticDocument {
	documents := make([]ElasticDocument, len(blog.Posts))
	for i, post := range blog.Posts {
		action := ElasticAction{Index: IndexMeta{IndexName: "myblog", Id: post.getId()}}
		fields := ElasticFields{Language: post.getLanguage(), Title: post.getTitle(), Body: post.getBody(), Url: post.getUrl(), Thumbnail: post.getThumbnail()}
		document := ElasticDocument{Action: action, Fields: fields}
		documents[i] = document
	}
	return documents
}

type ElasticEncoder struct {
	encoder			*json.Encoder
}

func NewElasticEncoder(w io.Writer) *ElasticEncoder {
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.SetEscapeHTML(false)
	return &ElasticEncoder{encoder: jsonEncoder}
}

func (enc *ElasticEncoder) Encode(blog BlogFeed) error {
	documents := newElasticDocuments(blog)
	for _, document := range documents {
		err := enc.encoder.Encode(document.Action)
        	if err != nil {
			return err
		}
		err = enc.encoder.Encode(document.Fields)
        	if err != nil {
			return err
		}
	}
	return nil
}
