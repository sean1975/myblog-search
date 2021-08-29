package convert


type VespaDocument struct {
	Put string		`json:"put"`
	Fields Fields		`json:"fields"`
}

type Fields struct {
	Language string		`json:"language"`
	Id string		`json:"id"`
	Title string		`json:"title"`
	Body string		`json:"body"`
	Url string		`json:"url"`
	Thumbnail string	`json:"thumbnail"`
}
