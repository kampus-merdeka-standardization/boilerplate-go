package product_usecase

import (
	"context"
	product_response "simple-golang-database/internal/modules/product/model/response"
)

func (productUsecase *productUsecase) GetAllProduct(ctx context.Context) ([]product_response.Product, error) {
	products, err := productUsecase.productRepository.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	var productsRes []product_response.Product
	for _, v := range products {
		productsRes = append(productsRes, product_response.Product{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
		})
	}

	return productsRes, nil
}
