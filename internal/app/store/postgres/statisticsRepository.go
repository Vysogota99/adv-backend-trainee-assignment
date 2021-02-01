package postgres

import (
	"fmt"
	"time"

	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/models"
)

// StatisticsRepository - postgres implementation of StatisticsRepository
type StatisticsRepository struct {
	store *Store
}

// Save - save statistics to store
func (s *StatisticsRepository) Save(stat *models.Statistics) error {
	tx, err := s.store.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	query := `
				INSERT INTO statisctics(date, views, clicks, cost)
				VALUES ($1, $2, $3, $4)
			`

	_, err = tx.Exec(query, stat.Date, stat.Views, stat.Clicks, stat.Cost)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

// Delete - delete all rows from talbe staistics
func (s *StatisticsRepository) Delete() (int, error) {
	tx, err := s.store.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return 0, err
	}

	query := `
				DELETE FROM statisctics
			`

	result, err := tx.Exec(query)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	tx.Commit()

	return int(rowsDeleted), nil
}

// GetInRange - ...
func (s *StatisticsRepository) GetInRange(from, to, orderBy string) ([]models.Statistics, error) {
	tx, err := s.store.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	query := `
				SELECT date, views, clicks, cost, cost/clicks as cpc, (cost/views)*1000 as cpm
				FROM statisctics
				WHERE date >= $1 AND date <= $2
				ORDER BY %s ASC 

				`
	if orderBy == "" {
		orderBy = "date"
	}

	query = fmt.Sprintf(query, orderBy)

	stmt, err := tx.Preparex(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(from, to)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.Statistics, 0)
	for rows.Next() {
		stat := models.Statistics{}
		var date string
		if err := rows.Scan(&date, &stat.Views, &stat.Clicks, &stat.Cost, &stat.Cpc, &stat.Cpm); err != nil {
			return nil, err
		}

		dateTime, err := time.Parse(time.RFC3339, date)
		if err != nil {
			return nil, err
		}

		stat.Date = dateTime
		result = append(result, stat)
	}

	tx.Commit()
	return result, nil
}
