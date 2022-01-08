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
		g.Middleware(middleware.CORS, middleware.Logger)
		g.POST("/auth/token", middleware.Auth.LoginHandler)
		g.GET("/auth/token", middleware.Auth.RefreshHandler)
		g.DELETE("/auth/token", middleware.Auth.LogoutHandler)
		g.POST("/user", api.User.SignUp)
		g.GET("/categories", api.Category.Get)
		g.GET("/products/{productId}", api.Product.GetProduct)
		g.GET("/products", api.Product.GetProductList)
	})
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.CORS, middleware.Logger, middleware.MiddlewareAuth)
		g.GET("/profile", api.User.Profile)
		g.PUT("/balance", api.User.AddBalance)
	})
	s.Group("/orders/", func(g *ghttp.RouterGroup) {
		g.Middleware(middleware.CORS, middleware.Logger, middleware.MiddlewareAuth)
		g.GET("/{orderId}", api.Order.GetOrder)
		g.GET("/", api.Order.GetOrderList)
		g.POST("/", api.User.BuyProduct)
	})
}
