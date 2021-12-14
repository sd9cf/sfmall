package service

import (
	"context"
	"errors"
	"sfmall/app/dao"
	"sfmall/app/model"
	"sfmall/library/paging"
)

var Product = productService{}

type productService struct{}


func(s *productService) GetProduct(id string) (*model.Product, error) {
	var product *model.Product
	err := dao.Product.Ctx(context.TODO()).Where(id).Scan(product)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("未找到商品")
	}
	return product, nil
}


func(s *productService) GetProducts(req *model.ProductApiGetProductsReq) ([]*model.SimpleProduct, *paging.Paging, error) {
	categoryId := req.CategoryId
	db := dao.Product.Ctx(context.TODO()).Where("category_id=?", categoryId).Order(
		req.OrderColumn + " " + req.OrderType)
	total, err := db.Count()
	p := paging.Create(req.PageNum, req.PageSize, total)
	db.Limit(p.PageSize, p.StartNum)
	if err !=nil {
		return nil, nil, err
	}
	product, err := db.All()
	if err !=nil {
		return nil, nil, err
	}
	if product == nil {
		return nil, nil, errors.New("未找到商品")
	}
	var productlist []*model.SimpleProduct
	if err = product.Structs(&productlist); err != nil {
		return nil, nil, err
	} 

	return productlist, p, nil
}