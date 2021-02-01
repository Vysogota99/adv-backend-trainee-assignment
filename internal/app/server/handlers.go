package server

import (
	"net/http"
	"time"

	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/models"
	"github.com/gin-gonic/gin"
)

// SaveStatHandler - save statistics to store
func (r *Router) SaveStatHandler(c *gin.Context) {

	type request struct {
		Date   string  `json:"date" binding:"required,valid_date"`
		Views  uint    `json:"views,omitempty"`
		Clicks uint    `json:"clicks,omitempty"`
		Cost   float64 `json:"cost,omitempty" binding:"min=0"`
	}

	req := request{}
	if err := c.ShouldBindJSON(&req); err != nil {
		respond(c, http.StatusBadRequest, "", err.Error())
		return
	}

	date, err := time.Parse(LAYOUT, req.Date)
	if err != nil {
		respond(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	stat := models.Statistics{
		Date:   date,
		Views:  req.Views,
		Clicks: req.Clicks,
		Cost:   req.Cost,
	}

	if err := r.store.Stat().Save(&stat); err != nil {
		respond(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	respond(c, http.StatusOK, "готово", "")
}

// DeleteHandler - delete all statistics from store
func (r *Router) DeleteHandler(c *gin.Context) {
	nDeleted, err := r.store.Stat().Delete()
	if err != nil {
		respond(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	respond(c, http.StatusOK, map[string]int{
		"rows deleted": nDeleted,
	}, "")
}

// GetStatHandler - return statistics in historical range (from, to)
func (r *Router) GetStatHandler(c *gin.Context) {
	type query struct {
		From    time.Time `form:"from" binding:"required" time_format:"2006-01-02"`
		To      time.Time `form:"to" binding:"required,gtfield=From" time_format:"2006-01-02"`
		OrderBy string    `form:"orderBy" binding:"omitempty,oneof=date views clicks cost cpc cpm"`
	}

	q := query{}
	if err := c.ShouldBindQuery(&q); err != nil {
		respond(c, http.StatusBadRequest, "", err.Error())
		return
	}

	fromStr := q.From.Format(LAYOUT)
	toStr := q.To.Format(LAYOUT)

	stats, err := r.store.Stat().GetInRange(fromStr, toStr, q.OrderBy)
	if err != nil {
		respond(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	respond(c, http.StatusOK, stats, "")
}

func respond(c *gin.Context, code int, result interface{}, err string) {
	c.JSON(
		code,
		gin.H{
			"result": result,
			"error":  err,
		},
	)
}
