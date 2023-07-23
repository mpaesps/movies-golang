package handler

import (
	"movies-crud/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UpdateMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var movie models.Movie
		if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado!"})
			return
		}

		var updatedMovie models.Movie
		if err := c.ShouldBindJSON(&updatedMovie); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Erro": err.Error()})
			return
		}

		db.Model(&movie).Updates(updatedMovie)
		c.JSON(http.StatusOK, gin.H{"data": movie})
	}
}
