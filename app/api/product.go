package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"sfmall/library/response"
	"sfmall/app/model"
	"sfmall/app/service"
)

var Product = new(productApi)

type productApi struct{}

func (a *productApi) GetProduct(r *ghttp.Request) {
	productId := r.Get("productId")
	product, err := service.Product.GetProduct(gconv.String(productId))
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", product)
}

func (a *productApi) GetProductList(r *ghttp.Request) {
	var apiReq *model.ProductApiGetProductsReq
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	products, err := service.Product.GetProducts(apiReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", products)
}