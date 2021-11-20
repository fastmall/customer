package service

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"github.com/fastmall/customer/api"
	"math/rand"
	"os"
	"time"
)

type CustomerService struct {
	api.UnimplementedCustomerServiceServer
}

func (c CustomerService) CreateCustomer(ctx context.Context, req *api.CreateCustomerRequest) (*api.CreateCustomerResponse, error) {
	logger.Infof("receive invoke :%v", req)
	rand.Seed(time.Now().UnixNano())
	hostname, _ := os.Hostname()
	res := api.CreateCustomerResponse{
		CustomerId: rand.Int63(),
		Token:      "from:" + hostname,
	}
	return &res, nil
}

func (c CustomerService) mustEmbedUnimplementedCustomerServiceServer() {
	panic("implement me")
}
