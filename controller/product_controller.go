package controller

import (
	"net/http"

	"go-api/model"
	usecase "go-api/useCase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {

	ProductUsecase usecase.ProductUsecase

}

func NewProductController(usecase usecase.ProductUsecase) ProductController {

	return ProductController{
		ProductUsecase: usecase,
	}

}

func (p *ProductController) GetProducts(ctx *gin.Context){



	products := []model.Product{
		{

			ID: 1,
			Name: "Batata",
			Price: 12.0,
		},


	}


	ctx.JSON(http.StatusOK, products)
}