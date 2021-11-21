package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/dubbogo/triple/pkg/triple"
	"github.com/fastmall/customer/service"
	goodsApi "github.com/fastmall/goods/api"
)

var GoodsService = &goodsApi.GoodsServiceClientImpl{}

func StartDubbo() {
	config.SetConsumerService(GoodsService)
	config.SetProviderService(&service.CustomerService{})
	err := config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
	err = config.Load(config.WithPath("conf/dubbo.yaml"))
	if err != nil {
		logger.Fatal(err)
		return
	}
}
