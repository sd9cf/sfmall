package model

type OrderApiGetOrdersReq struct {
	UserId string `json:"userId"`
	PagingQueryReq
}