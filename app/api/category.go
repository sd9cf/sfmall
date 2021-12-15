package api

import (
	"github.com/gogf/gf/net/ghttp"
	"sfmall/library/response"
	"sfmall/app/service"
)

var Category = new(categoryApi)

type categoryApi struct{}

func (a *categoryApi) Get(r *ghttp.Request) {
	categories, err := service.Category.GetCategory()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", categories)
}