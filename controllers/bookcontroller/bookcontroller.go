package bookcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haviz000/rest-api-gin/database"
	"github.com/haviz000/rest-api-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var books []models.Book
	DB := database.GetDB()
	DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func Show(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	DB := database.GetDB()
	if err := DB.First(&book, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"messages": "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"messages": err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func Create(c *gin.Context) {
	var book models.Book
	DB := database.GetDB()
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	DB.Create(&book)
	c.String(http.StatusOK, "created")
}

func Update(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	DB := database.GetDB()

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"messages": err.Error(),
		})
		return
	}
	if DB.Model(&book).Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"messages": "tidak dapat mengupdate produk",
		})
		return
	}
	c.String(http.StatusOK, "updated")
}

func Delete(c *gin.Context) {

	var book models.Book
	DB := database.GetDB()
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if DB.Delete(&book, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus book"})
		return
	}

	c.String(http.StatusOK, "deleted")
}
