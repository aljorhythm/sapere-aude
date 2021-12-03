package collection

type CollectionItem struct {
	Id   string
	Url  string
	Tags []*Tag
}

func NewCollectionItem(id string, url string, tags []*Tag) *CollectionItem {
	return &CollectionItem{
		id, url, tags,
	}
}
