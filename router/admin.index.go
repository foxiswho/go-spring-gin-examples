package router

import (
	"github.com/go-spring/spring-core/gs"
	"github/foxiswho/go-spring-gin-examples/app/domainDemo/controller"
	"github/foxiswho/go-spring-gin-examples/middleware/server/ginServer"
)

func init() {
	// 首页路由
	gs.Object(&controller.Index{}).Init(func(s *controller.Index) {
		r := ginServer.GinServerDefault
		r.GET("/", s.Index)
		r.GET("/port", s.Port)
	})
}
