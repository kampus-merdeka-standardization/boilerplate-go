package product_grpc

import (
	"context"

	product "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/grpc"
)

func (p *productServer) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.ProductResponse, error) {
	panic("implement me")
}

func (p *productServer) GetAllProduct(ctx context.Context, req *product.Empty) (*product.GetAllProductResponse, error) {
	panic("implement me")
}

func (p *productServer) GetProductByID(ctx context.Context, req *product.ID) (*product.ProductResponse, error) {
	panic("implement me")
}
func (p *productServer) UpdateProductByID(ctx context.Context, req *product.UpdateProductByIDRequest) (*product.ProductResponse, error) {
	panic("implement me")
}
func (p *productServer) DeleteProductByID(ctx context.Context, req *product.ID) (*product.ID, error) {
	panic("implement me")
}
