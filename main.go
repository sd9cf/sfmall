package main

import (
	_ "sfmall/boot"
	_ "sfmall/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
