package mock

import (
	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/models"
)

// StatisticsRepository - postgres implementation of StatisticsRepository
type StatisticsRepository struct {
	store *Store
}

// Save - save statistics to store
func (s *StatisticsRepository) Save(stat *models.Statistics) error {
	return nil
}

// Delete - delete all rows from talbe staistics
func (s *StatisticsRepository) Delete() (int, error) {
	return 0, nil
}

// GetInRange - ...
func (s *StatisticsRepository) GetInRange(from, to, orderBy string) ([]models.Statistics, error) {
	res := make([]models.Statistics, 2)
	return res, nil
}
