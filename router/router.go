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
		g.Middleware(middleware.CORS)
		g.POST("/login", middleware.Auth.LoginHandler)
		g.GET("/refresh_token", middleware.Auth.RefreshHandler)
		g.GET("/logout", middleware.Auth.LogoutHandler)
		g.POST("/signup", api.User.SignUp)
		g.GET("/category", api.Category.Get)
		g.GET("/product", api.Product.GetProduct)
		g.GET("/products", api.Product.GetProductList)
	})
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.CORS, middleware.MiddlewareAuth)
		g.GET("/profile", api.User.Profile)
		g.PUT("/balance", api.User.AddBalance)
		g.GET("/order", api.Order.GetOrder)
		g.GET("/orders", api.Order.GetOrderList)
		g.POST("/order", api.User.BuyProduct)
	})
}
