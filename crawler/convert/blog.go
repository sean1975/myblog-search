package convert

import (
	"encoding/xml"
	"github.com/grokify/html-strip-tags-go"
	"html"
	"strings"
)

type BlogFeed struct {
	Posts []Post		`xml:"entry"`
}

type Post struct {
	Id Id			`xml:"id"`
	Updated string		`xml:"updated"`
	Categories []Category	`xml:"category"`
	Title string		`xml:"title"`
	Body Body		`xml:"content"`
	Link Link		`xml:"link"`
	Thumbnail Thumbnail	`xml:"thumbnail"`
}

type Id string

func (a *Id) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	id := strings.Split(s, "post-")
	*a = Id(id[len(id)-1])
	return nil
}

func getAttributeValue(element xml.StartElement, name string) string {
	for _, attr := range element.Attr {
		if attr.Name.Local == name {
			return attr.Value
		}
        }
	return ""
}

type Category string

func (a *Category) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	value := getAttributeValue(start, "term")
	if len(value) > 0 {
		*a = Category(value)
	}
	return nil
}

type Body string

func trimSpace(s string) string {
	unescaped := html.UnescapeString(s)
	return strings.Join(strings.Fields(strings.TrimSpace(unescaped)), " ")
}

func (a *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	trimmed := trimSpace(strip.StripTags(s))
	trimmed = strings.TrimPrefix(trimmed, "English version ")
	*a = Body(trimmed)
	return nil
}

type Link string

func (a *Link) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	value := getAttributeValue(start, "rel")
	if value != "alternate" {
		return nil
	}
	value = getAttributeValue(start, "href")
	if len(value) > 0 {
		*a = Link(value)
	}
	return nil
}

type Thumbnail string

func (a *Thumbnail) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	value := getAttributeValue(start, "url")
	if len(value) > 0 {
		*a = Thumbnail(value)
	}
	return nil
}

func (a *Post) getId() string {
	return string(a.Id)
}

func (a *Post) getLanguage() string {
	for _, category := range a.Categories { 
		if string(category) == "Australia" {
			return "en"
		}
	}
	return "zh-TW"
}

func (a *Post) getTitle() string {
	return a.Title
}

func (a *Post) getBody() string {
	return string(a.Body)
}

func (a *Post) getUrl() string {
	return string(a.Link)
}

func (a *Post) getThumbnail() string {
	return string(a.Thumbnail)
}
