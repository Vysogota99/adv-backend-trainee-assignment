package store

import (
	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/models"
)

// Store - интерфейс для работы с хранилищем
type Store interface {
	Stat() StatisticsRepository
}

// StatisticsRepository - интерфейс, содержащий методы для получения статистики
type StatisticsRepository interface {
	Save(stat *models.Statistics) error
	Delete() (int, error)
	GetInRange(from, to, orderBy string) ([]models.Statistics, error)
}
