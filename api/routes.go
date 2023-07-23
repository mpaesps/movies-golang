package api

import (
	"movies-crud/internal/handler"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/movies", handler.CreateMovie(db))
	r.GET("/movies/:id", handler.GetMovie(db))
	r.PUT("/movies/:id", handler.UpdateMovie(db))
	r.DELETE("/movies/:id", handler.DeleteMovie(db))
}
