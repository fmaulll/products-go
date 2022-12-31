package main

import (
	"github.com/fmaulll/products-go/controllers/productscontroller"
	"github.com/fmaulll/products-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.GET("/api/products", productscontroller.Index)
	router.GET("/api/product/:id", productscontroller.Show)
	router.POST("/api/product", productscontroller.Create)
	router.PATCH("/api/product/:id", productscontroller.Update)
	router.DELETE("/api/product", productscontroller.Delete)

	router.Run()
}
