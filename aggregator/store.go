package main

import "github.com/olzh2102/toll-calculator/types"

type MemomyStore struct {
}

func NewMemoryStore() *MemomyStore {
	return &MemomyStore{}
}

func (s *MemomyStore) Insert(d types.Distance) error {
	return nil
}
