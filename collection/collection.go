package collection

type collection struct {
	articles []*Article
	userId   string
}

func (c *collection) GetUserId() string {
	return c.userId
}

func (c *collection) AddArticle(article *Article) error {
	c.articles = append(c.articles, article)
	return nil
}

func (c *collection) GetArticles() ([]*Article, error) {
	return c.articles, nil
}

type Collection interface {
	GetArticles() ([]*Article, error)
	AddArticle(*Article) error
	GetUserId() string
}

func NewCollection(userId string) Collection {
	return &collection{
		userId: userId,
	}
}
