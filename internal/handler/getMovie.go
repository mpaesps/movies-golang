package handler

import (
	"movies-crud/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetMovie(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var movie models.Movie
		if err := db.Where("id = ?", id).First(&movie).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"Erro": "Registro não encontrado!"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": movie})
	}
}
