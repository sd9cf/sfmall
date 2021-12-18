package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"

	"sfmall/library/response"
	"sfmall/app/model"
	"sfmall/app/service"
)

var User = new(userApi)

type userApi struct{}

func (a *userApi) SignUp(r *ghttp.Request) {
	var (
		apiReq     *model.UserApiSignUpReq
		serviceReq *model.UserServiceSignUpReq
	)
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.User.SignUp(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	} else {
		response.JsonExit(r, 0, "ok")
	}
}

func (a *userApi) Profile(r *ghttp.Request) {
	var id = r.Get("id")
	var profile *model.UserProfile
	var err error
	if profile, err = service.User.GetProfile(gconv.String(id)); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", profile)
}

func (a *userApi) AddBalance(r *ghttp.Request) {
	var apiReq *model.AddBalance
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	var id = r.Get("id")
	err := service.User.AddBalance(gconv.String(id), apiReq.Money)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}

func (a *userApi) BuyProduct(r *ghttp.Request) {
	var apiReq *model.BuyProduct
	var id = r.Get("id")
	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	err := service.User.BuyProduct(gconv.String(id), apiReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok")
}