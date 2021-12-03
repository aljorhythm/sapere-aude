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
		article := NewCollectionItem(
			"some.place.in.the.world",
			"hello",
			[]*Tag{},
		)
		err := collection.AddCollectionItem(article)

		assert.NoError(t, err)

		t.Run("Get collectionItems in collection", func(t *testing.T) {
			articles, err := collection.GetCollectionItems()

			wanted := []*CollectionItem{
				article,
			}
			assert.NoError(t, err)
			assert.Equal(t, wanted, articles)
		})
	})
}
