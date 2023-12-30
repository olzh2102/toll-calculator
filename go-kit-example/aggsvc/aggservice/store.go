package aggservice

import (
	"fmt"

	"github.com/olzh2102/toll-calculator/types"
)

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

func (s *MemomyStore) Get(id int) (float64, error) {
	dist, ok := s.data[id]
	if !ok {
		return 0.0, fmt.Errorf("could not find distance for obu id %d", id)
	}
	return dist, nil
}
