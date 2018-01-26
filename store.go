package cachestore

import (
	"errors"

	"github.com/fsm/fsm"
)

// New returns an instance of a cacheStore
func New() fsm.Store {
	return &cacheStore{
		traversers: make(map[string]fsm.Traverser, 0),
	}
}

type cacheStore struct {
	traversers map[string]fsm.Traverser
}

func (s *cacheStore) FetchTraverser(uuid string) (fsm.Traverser, error) {
	if traverser, ok := s.traversers[uuid]; ok {
		return traverser, nil
	}
	return nil, errors.New("Traverser does not exist")
}

func (s *cacheStore) CreateTraverser(uuid string) (fsm.Traverser, error) {
	if _, ok := s.traversers[uuid]; ok {
		return nil, errors.New("Traverser with UUID already exists")
	}
	traverser := &cachedTraverser{
		Data: make(map[string]interface{}, 0),
	}
	traverser.SetUUID(uuid)
	s.traversers[uuid] = traverser
	return traverser, nil
}
