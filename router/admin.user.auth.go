package router

import (
	"github.com/go-spring/spring-core/gs"
	"github/foxiswho/go-spring-gin-examples/app/domainDemo/controller"
	"github/foxiswho/go-spring-gin-examples/middleware/auth"
	"github/foxiswho/go-spring-gin-examples/middleware/server/ginServer"
)

func init() {
	// 用户
	gs.Object(&controller.User{}).Init(func(s *controller.User) {
		r := ginServer.GinServerDefault
		group := r.Group("/admin", auth.GroupMiddlewareOne(s.Sp))
		group.POST("/user/create", s.Create)
	})
	//数据库保存
	gs.Object(&controller.UserDb{}).Init(func(s *controller.UserDb) {
		r := ginServer.GinServerDefault
		//
		r.GET("/db-config", s.DbConfig)
		//
		group := r.Group("/admin", auth.GroupMiddlewareOne(s.Sp))
		group.POST("/user-db/create", s.Save)

	})
}
