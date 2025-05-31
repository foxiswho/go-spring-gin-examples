package controller

import (
	"github.com/gin-gonic/gin"
	"github/foxiswho/go-spring-gin-examples/app/domainDemo/model/modUser"
	"github/foxiswho/go-spring-gin-examples/app/domainDemo/service"
	"github/foxiswho/go-spring-gin-examples/middleware/auth"
	"github/foxiswho/go-spring-gin-examples/middleware/validator"
	"github/foxiswho/go-spring-gin-examples/pkg/config"
	"github/foxiswho/go-spring-gin-examples/pkg/tools/wrapperPg/rg"
)

// 用户数据 存入 db
type UserDb struct {
	sv       *service.UserDbService  `autowire:""`
	Sp       *auth.MiddlewareGroupSp `autowire:""`
	database config.Database         `value:"${database}"`
}

// 保存
func (t *UserDb) Save(c *gin.Context) {
	var ct modUser.CreateUpdateCt
	// 结构体数据绑定
	err := c.ShouldBind(&ct)
	if err != nil {
		//对 返回 错误进行转义 成中文
		translate := validator.Translate(err, &ct)
		if len(translate) > 0 {
			c.JSON(500, rg.ErrorMessageData[string](translate))
			return
		}
		c.JSON(500, rg.ErrorDefault[string]())
		return
	}
	c.JSON(200, t.sv.Save(c, ct))
}

// 输出 配置
func (t *UserDb) DbConfig(c *gin.Context) {
	c.JSON(200, rg.OkData(t.database))
}
