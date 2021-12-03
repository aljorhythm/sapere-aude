package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection(t *testing.T) {
	userId := "aljorhythm"
	collection := NewCollection(userId)

	assert.Equal(t, userId, collection.GetUserId())
	t.Run("Add new article to collection", func(t *testing.T) {
		article := NewArticle(
			"some.place.in.the.world",
			"hello",
			[]*Tag{},
		)
		err := collection.AddArticle(article)

		assert.NoError(t, err)

		t.Run("Get articles in collection", func(t *testing.T) {
			articles, err := collection.GetArticles()

			wanted := []*Article{
				article,
			}
			assert.NoError(t, err)
			assert.Equal(t, wanted, articles)
		})
	})
}
