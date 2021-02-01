package postgres

import (
	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/store"
	"github.com/jmoiron/sqlx"
)

// Store - postgres implementation of store
type Store struct {
	DB *sqlx.DB
	statisticsRepository store.StatisticsRepository
}

// New - helper to init Store
func New(db *sqlx.DB) *Store {
	return &Store{
		DB: db,
	}
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



