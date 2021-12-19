package model

type OrderApiGetOrderReq struct {
	OrderId string `v:"required#订单ID不为空"`
}

type OrderApiGetOrdersReq struct {
	PagingQueryReq
}