package postgres

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

const LAYOUT = "2006-01-02"

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := `
				INSERT INTO statisctics(date, views, clicks, cost)
				VALUES ($1, $2, $3, $4)
			`
	date := time.Now()
	stat := models.Statistics{
		Date:   date,
		Views:  1,
		Clicks: 1,
		Cost:   1,
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(stat.Date, stat.Views, stat.Clicks, stat.Cost).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	store := New(sqlxDB)

	err = store.Stat().Save(&stat)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := `
				DELETE FROM statisctics
			`

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	store := New(sqlxDB)

	n, err := store.Stat().Delete()
	assert.NoError(t, err)
	assert.Equal(t, 0, n)
}

func TestGetInRange(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	query := `
				SELECT date, views, clicks, cost, cost/clicks as cpc, (cost/views)*1000 as cpm
				FROM statisctics
				WHERE date >= $1 and date <= $2
				ORDER BY date ASC 
			`

	from := "2010-02-01"
	to := "2010-03-03"
	rows := mock.NewRows(
		[]string{"date", "views", "clicks", "cost", "cpc", "cpm"},
	).AddRow("2010-03-01", 10, 4, 0.01, 0.00250000000000000000, 1.00000000000000000000).
		AddRow("2010-03-01", 90, 40, 0.01, 0.00025000000000000000, 0.11111111111111111000).
		AddRow("2010-03-01", 90, 40, 0.01, 0.00025000000000000000, 0.11111111111111111000)

	mock.ExpectBegin()
	stmt := mock.ExpectPrepare(regexp.QuoteMeta(query))
	stmt.ExpectQuery().WithArgs(from, to).WillReturnRows(rows)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	store := New(sqlxDB)

	_, err = store.Stat().GetInRange(from, to, "")
	assert.NoError(t, err)
}
