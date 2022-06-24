package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"lebron/apps/app/api/internal/config"
	"lebron/apps/order/rpc/order/order"
	"lebron/apps/product/rpc/product/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
