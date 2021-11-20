package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/customer/api"
	"math/rand"
	"time"
)

type CustomerService struct {
	api.UnimplementedCustomerServiceServer
}

func (c CustomerService) CreateCustomer(ctx context.Context, req *api.CreateCustomerRequest) (*api.CreateCustomerResponse, error) {
	logger.Info("receive invoke :%v", req)
	rand.Seed(time.Now().UnixNano())
	res := api.CreateCustomerResponse{
		UserId: rand.Int63(),
	}
	return &res, nil
}

func (c CustomerService) mustEmbedUnimplementedCustomerServiceServer() {
	panic("implement me")
}
