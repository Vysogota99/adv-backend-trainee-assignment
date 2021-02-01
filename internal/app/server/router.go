package server

import (
	"fmt"
	"time"

	"github.com/Vysogota99/adv-backend-trainee-assignment/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

const LAYOUT = "2006-01-02"

// Router ...
type Router struct {
	router     *gin.Engine
	serverPort string
	validator  *validator.Validate
	store      store.Store
}

// NewRouter - helper for initialization http router
func NewRouter(serverPort string, store store.Store) *Router {
	return &Router{
		router:     gin.Default(),
		serverPort: serverPort,
		store:      store,
	}
}

// Setup - setup validator and routes in router
func (r *Router) Setup() (*gin.Engine, error) {
	validator, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil, fmt.Errorf("Cant init validator")
	}

	if err := validator.RegisterValidation("valid_date", validateDate); err != nil {
		return nil, err
	}

	r.validator = validator
	r.router.POST("/stat", r.SaveStatHandler)
	r.router.DELETE("/stat", r.DeleteHandler)
	r.router.GET("/stat", r.GetStatHandler)

	return r.router, nil
}

func validateDate(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(string); ok {
		_, err := time.Parse(LAYOUT, date)
		if err != nil {
			return false
		}

		return true
	}

	return false
}
