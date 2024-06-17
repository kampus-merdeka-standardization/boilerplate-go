package product_grpc

import (
	product "simple-golang-database/internal/modules/product/grpc"
)

type productServer struct {
	product.UnimplementedProductServiveServer
}
