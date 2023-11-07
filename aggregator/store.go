package main

import "github.com/olzh2102/toll-calculator/types"

type MemomyStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemomyStore {
	return &MemomyStore{
		data: make(map[int]float64),
	}
}

func (s *MemomyStore) Insert(d types.Distance) error {
	s.data[d.OBUID] += d.Value
	return nil
}
