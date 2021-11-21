package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/dubbogo/triple/pkg/triple"
	cartApi "github.com/fastmall/cart/api"
	"github.com/fastmall/customer/service"
	goodsApi "github.com/fastmall/goods/api"
	orderApi "github.com/fastmall/order/api"
)

var GoodsService = &goodsApi.GoodsServiceClientImpl{}
var CartService = &cartApi.CartServiceClientImpl{}
var OrderService = &orderApi.OrderServiceClientImpl{}

func StartDubbo() {
	config.SetConsumerService(GoodsService)
	config.SetConsumerService(CartService)
	config.SetConsumerService(OrderService)
	config.SetProviderService(&service.CustomerService{})
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
