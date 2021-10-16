package convert

type IndexEncoder interface {
	Encode(blog BlogFeed) error
}
