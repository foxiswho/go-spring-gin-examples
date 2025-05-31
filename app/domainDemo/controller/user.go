package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/util/syslog"
	"github/foxiswho/go-spring-gin-examples/app/domainDemo/model/modUser"
	"github/foxiswho/go-spring-gin-examples/middleware/auth"
	"github/foxiswho/go-spring-gin-examples/middleware/validator"
	"github/foxiswho/go-spring-gin-examples/pkg/tools/wrapperPg/rg"
)

// 用户
type User struct {
	Sp *auth.MiddlewareGroupSp `autowire:""`
}

// 测试 post 结构体 json 验证
func (t User) Create(c *gin.Context) {
	var ct modUser.CreateUpdateCt
	//
	err := c.ShouldBind(&ct)
	if err != nil {
		translate := validator.Translate(err, &ct)
		if len(translate) > 0 {
			c.JSON(500, rg.ErrorMessageData[string](translate))
			return
		}
		c.JSON(500, rg.ErrorDefault[string]())
		return
	}
	syslog.Infof("ct: %+v", ct)
	//
	c.JSON(200, rg.OkDefault[string]())
}
