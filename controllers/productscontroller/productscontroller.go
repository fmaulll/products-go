package productscontroller

import (
	"encoding/json"
	"net/http"

	"github.com/fmaulll/products-go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(context *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)

	context.JSON(http.StatusOK, gin.H{"products": products})
}

func Show(context *gin.Context) {
	var product models.Product

	id := context.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
			return

		default:
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(context *gin.Context) {
	var product models.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	context.JSON(http.StatusCreated, gin.H{"product": product})
}

func Update(context *gin.Context) {
	var product models.Product

	id := context.Param("id")

	if err := context.ShouldBindJSON(&product); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Product updated"})
}

func Delete(context *gin.Context) {
	var product models.Product

	var input struct {
		Id json.Number
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't delete product!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
