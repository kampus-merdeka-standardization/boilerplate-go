package product_grpc

import (
	product "github.com/kampus-merdeka-standardization/boilerplate-go/pkg/proto/gen/product/v1"
)

type productServer struct {
	product.UnimplementedProductServiveServer
}
