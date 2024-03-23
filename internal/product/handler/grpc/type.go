package product_grpc

import (
	product "github.com/kampus-merdeka-standardization/boilerplate-go/internal/product/generated/grpc"
)

type productServer struct {
	product.UnimplementedProductServiveServer
}
