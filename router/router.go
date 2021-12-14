package router

import (
	"sfmall/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"sfmall/app/middleware"
)


func init() {
	s := g.Server()
	s.Group("/", func(g *ghttp.RouterGroup) {
		g.ALL("/hello", api.Hello)
		g.ALL("/login", middleware.Auth.LoginHandler)
		g.ALL("/refresh_token", middleware.Auth.RefreshHandler)
		g.ALL("/logout", middleware.Auth.LogoutHandler)
		g.ALL("/signup", api.User.SignUp)
	})
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.CORS, middleware.MiddlewareAuth)
		g.ALL("/profile", api.User.Profile)
	})
}
