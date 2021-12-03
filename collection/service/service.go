package service

import "github.com/aljorhythm/sapere-aude/collection"

type Service interface {
	AddArticleToCollection(article *collection.CollectionItem, collection *collection.Collection) error
	AddTagToArticle(collection *collection.Collection, article *collection.CollectionItem) error
}
