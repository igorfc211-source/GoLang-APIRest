package controller

import (
	"net/http"
	"strconv"

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

func (p *ProductController) GetProducts(ctx *gin.Context) {

	products, err := p.ProductUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func(p *ProductController) CreateProduct(ctx *gin.Context){

	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {

		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.ProductUsecase.CreateProduct(product)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (p *ProductController) GetProductsById(ctx *gin.Context) {

	id := ctx.Param("productId")

	if id == ""{

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "productId não pode ser vazio",
		})
		return
	}
	
	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}


	products, err := p.ProductUsecase.GetProductByID(productId)
	if err != nil {

		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

		if products == nil {
			response := model.Response{
				Message: "produto não foi encontrado",
			}
			ctx.JSON(http.StatusNotFound, response)
			return

		}
	ctx.JSON(http.StatusOK, products)
}