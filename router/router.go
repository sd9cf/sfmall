package router

import (
	"sfmall/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"sfmall/app/service"
	"sfmall/app/api"
)

func middlewareAuth(r *ghttp.Request) {
	api.Auth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func init() {
	s := g.Server()
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.ALL("/hello", api.Hello)
		g.ALL("/login", api.Auth.LoginHandler)
		g.ALL("/refresh_token", api.Auth.RefreshHandler)
		g.ALL("/logout", api.Auth.LogoutHandler)
		g.ALL("/signup", api.User.SignUp)
	})
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Middleware.CORS, middlewareAuth)
		g.ALL("/profile", api.User.)
	})
}
