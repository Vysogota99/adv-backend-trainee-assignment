package models

import "time"

// Statistics ...
type Statistics struct {
	Date   time.Time `json:"date" binding:"required" db:"date"`
	Views  uint      `json:"views,omitempty" db:"views"`
	Clicks uint      `json:"clicks,omitempty" db:"clicks"`
	Cost   float64   `json:"cost,omitempty" db:"cost"`
	Cpc    float64   `json:"cpc,omitempty" db:"cpc"`
	Cpm    float64   `json:"cpm,omitempty" db:"cpm"`
}
