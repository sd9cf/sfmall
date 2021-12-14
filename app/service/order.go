package service

import (
	"context"
	"errors"
	"sfmall/app/dao"
	"sfmall/app/model"
	"sfmall/library/paging"
)

var Order = orderService{}

type orderService struct{}

func(s *orderService) GetOrders(req *model.OrderApiGetOrdersReq) ([]*model.Order, *paging.Paging, error) {
	userId := req.UserId
	db := dao.Order.Ctx(context.TODO()).Where("user_id=?", userId).Order(
		req.OrderColumn + " " + req.OrderType)
	total, err := db.Count()
	p := paging.Create(req.PageNum, req.PageSize, total)
	db.Limit(p.PageSize, p.StartNum)
	if err !=nil {
		return nil, nil, err
	}
	order, err := db.All()
	if err != nil {
		return nil, nil, err
	}
	if order == nil {
		return nil, nil, errors.New("未找到订单")
	}
	var orderlist []*model.Order
	if err = order.Structs(&orderlist); err != nil {
		return nil, nil, err
	} 

	return orderlist, p, nil
}

func (s *orderService) GetOrder(id string, userId string) (*model.Order, error) {
	var order *model.Order
	err := dao.Order.Ctx(context.TODO()).Where(id).Scan(order)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errors.New("未找到订单")
	}
	if order.UserId != userId {
		return nil, errors.New("未找到订单")
	}
	return order, nil
}