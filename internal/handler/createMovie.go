package handler

import (
	"movies-crud/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var movie models.Movie
		if err := c.ShouldBindJSON(&movie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
			return
		}

		if err := db.Create(&movie).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Erro": err.Error()})
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}
