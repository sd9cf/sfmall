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

var Product = productService{}

type productService struct{}


func(s *productService) GetProduct(id string) (*model.Product, error) {
	var product *model.Product
	err := dao.Product.Ctx(context.TODO()).WherePri(id).Scan(&product)
	if err != nil {
		g.Log().Error("数据库请求商品错误id:%s,错误%v", id, err)
		return nil, myerror.DATABASEERROR
	}
	if product == nil {
		return nil, errors.New("未找到商品")
	}
	return product, nil
}


func(s *productService) GetProducts(req *model.ProductApiGetProductsReq) (*model.PagingRes, error) {
	// categoryId := req.CategoryId
	// db := dao.Product.Ctx(context.TODO()).Where("category_id=?", categoryId).Order(
	// 	req.OrderColumn + " " + req.OrderType)
	db := dao.Product.Ctx(context.TODO()).Order(
		req.OrderColumn + " " + req.OrderType)
	total, err := db.Count()
	p := paging.Create(req.PageNum, req.PageSize, total)
	db.Limit(p.PageSize, p.StartNum)
	if err !=nil {
		g.Log().Error("数据库请求商品错误请求:%v,错误%v", req, err)
		return nil, myerror.DATABASEERROR
	}
	product, err := db.All()
	if err !=nil || product == nil {
		g.Log().Error("数据库请求商品错误请求:%v,错误%v", req, err)
		return nil, myerror.DATABASEERROR
	}
	if product == nil {
		return nil, errors.New("未找到商品")
	}
	var productlist []*model.SimpleProduct
	if err = product.Structs(&productlist); err != nil {
		g.Log().Errorf("数据%v映射错误", product)
		return nil, myerror.MAPPINGERROR
	} 
	
	return &model.PagingRes{
		Data: productlist,
		Paging: p,
	}, nil
}