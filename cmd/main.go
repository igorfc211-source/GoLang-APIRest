package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	usecase "go-api/useCase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil{	

		panic(err)

	}
	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada do usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	//camada dos controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{

			"message": "pong",
		})

	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductsById)
	server.Run(":8000")
}
