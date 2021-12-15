package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"sfmall/library/response"
	"sfmall/app/model"
	"sfmall/app/service"
)

var Order = new(orderApi)

type orderApi struct{}

func (a *orderApi) GetOrder(r *ghttp.Request) {
	var apiReq *model.OrderApiGetOrderReq
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	userId := r.Get("id")
	order, err := service.Order.GetOrder(apiReq.OrderId, gconv.String(userId))
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", order)
}

func (a *orderApi) GetOrderList(r *ghttp.Request) {
	var apiReq *model.OrderApiGetOrdersReq
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	userId := r.Get("id")
	orders, err := service.Order.GetOrders(apiReq, gconv.String(userId))
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", orders)
}