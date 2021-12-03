package service

import "github.com/aljorhythm/sapere-aude/entry/store"

type Servicer interface {
}

type Service struct {
}

func NewService(store store.Store) Servicer {
	return &Service{}
}
