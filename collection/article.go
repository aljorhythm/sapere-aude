package collection

type Article struct {
	Id   string
	Url  string
	Tags []*Tag
}

func NewArticle(id string, url string, tags []*Tag) *Article {
	return &Article{
		id, url, tags,
	}
}
