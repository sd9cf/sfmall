package model

type OrderApiGetOrderReq struct {
	OrderId string `json:"orderId"`
}

type OrderApiGetOrdersReq struct {
	PagingQueryReq
}