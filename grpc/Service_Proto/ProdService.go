package grpc_Service_Proto

import (
	"context"
)

type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, request *ProdRequest) (*ProdResponse, error) {{}
	return &ProdResponse{Stock:20}, nil
}
