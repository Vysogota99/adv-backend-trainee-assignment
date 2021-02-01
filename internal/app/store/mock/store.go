package mock

import (
	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/store"
)

// Store - postgres implementation of store
type Store struct {
	statisticsRepository store.StatisticsRepository
}

// New - helper to init Store
func New() *Store {
	return &Store{}
}

// Stat - implementation of StatisticsRepository interface
func (s *Store) Stat() store.StatisticsRepository {
	if s.statisticsRepository == nil {
		s.statisticsRepository = &StatisticsRepository{
			store: s,
		}
	}

	return s.statisticsRepository
}
