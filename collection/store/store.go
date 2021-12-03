package store

import "github.com/aljorhythm/sapere-aude/collection"

type Store interface {
	StoreCollection(collection *collection.Collection) error
	RetrieveCollection(userId string) error
}