package convert

import (
	"encoding/xml"
	"io"
)

func NewIndexEncoder(w io.Writer, documentType string) IndexEncoder {
	if (documentType == "vespa") {
		return NewVespaEncoder(w)
	} else if (documentType == "elastic") {
		return NewElasticEncoder(w)
	} else {
		return nil
	}
}

func XmlToJson(r io.Reader, w io.Writer, documentType string) error {
	decoder := xml.NewDecoder(r)
	blogFeed := BlogFeed{}
	err := decoder.Decode(&blogFeed)
        if err != nil {
		return err
	}
	encoder := NewIndexEncoder(w, documentType)
	err = encoder.Encode(blogFeed)
        if err != nil {
		return err
	}
	return nil
}
