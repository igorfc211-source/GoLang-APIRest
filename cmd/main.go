package main

import (
	"go-api/controller"
	"go-api/db"
	usecase "go-api/useCase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil{	

		panic(err)

	}
	ProductUsecase := usecase.NewProductUsecase()

	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{

			"message": "pong",
		})

	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
