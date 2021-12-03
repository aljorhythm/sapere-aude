package collection

type collection struct {
	collectionItems []*CollectionItem
	userId          string
}

func (c *collection) GetUserId() string {
	return c.userId
}

func (c *collection) AddCollectionItem(article *CollectionItem) error {
	c.collectionItems = append(c.collectionItems, article)
	return nil
}

func (c *collection) GetCollectionItems() ([]*CollectionItem, error) {
	return c.collectionItems, nil
}

type Collection interface {
	GetCollectionItems() ([]*CollectionItem, error)
	AddCollectionItem(*CollectionItem) error
	GetUserId() string
}

func NewCollection(userId string) Collection {
	return &collection{
		userId: userId,
	}
}
