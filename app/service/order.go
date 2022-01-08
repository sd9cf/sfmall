package service

import (
	"context"
	"errors"
	"github.com/gogf/gf/frame/g"
	"sfmall/app/dao"
	"sfmall/app/model"
	"sfmall/app/myerror"
	"sfmall/library/paging"
)

var Order = orderService{}

type orderService struct{}

func(s *orderService) GetOrders(req *model.OrderApiGetOrdersReq, userId string) (*model.PagingRes, error) {
	db := dao.Order.Ctx(context.TODO()).Where("user_id=?", userId).Order(
		req.OrderColumn + " " + req.OrderType)
	total, err := db.Count()
	p := paging.Create(req.PageNum, req.PageSize, total)
	db.Limit(p.PageSize, p.StartNum)
	if err !=nil {
		g.Log().Error("数据库错误请求:%v,错误%v", req, err)
		return nil, myerror.DATABASEERROR
	}
	order, err := db.All()
	if err != nil {
		g.Log().Error("数据库错误请求:%v,错误%v", req, err)
		return nil, myerror.DATABASEERROR
	}
	if order == nil {
		return nil, errors.New("未找到订单")
	}
	var orderlist []*model.Order
	if err = order.Structs(&orderlist); err != nil {
		g.Log().Errorf("数据%v映射错误", order)
		return nil, myerror.MAPPINGERROR
	} 

	return &model.PagingRes{
		Data: orderlist,
		Paging: p,
	}, nil
}

func (s *orderService) GetOrder(id string, userId string) (*model.Order, error) {
	var order *model.Order
	err := dao.Order.Ctx(context.TODO()).Where(id).Scan(&order)
	if err != nil {
		g.Log().Error("获取订单错误id:%s,错误%v", id, err)
		return nil, myerror.DATABASEERROR
	}
	if order == nil {
		return nil, errors.New("未找到订单")
	}
	if order.UserId != userId {
		return nil, errors.New("未找到订单")
	}
	return order, nil
}