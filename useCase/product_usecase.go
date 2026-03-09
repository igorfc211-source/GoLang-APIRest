package usecase

import model "go-api/model"

type ProductUsecase struct {
	//repository
}

func NewProductUsecase() ProductUsecase {

	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error){


	return []model.Product{}, nil
}