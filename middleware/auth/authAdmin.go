package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github/foxiswho/go-spring-gin-examples/pkg/tools/wrapperPg/rg"
	"time"
)

func init() {
	gs.Object(&MiddlewareGroupSp{})
}

// 中间件 服务
type MiddlewareGroupSp struct {
}

// 权限验证 中间件
func GroupMiddlewareOne(m *MiddlewareGroupSp) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("【admin】权限验证 中间件。time: " + time.Now().String())
		// 验证权限,url 参数 带 pass=ok，时，表示通过
		pass := c.Query("pass")
		if pass == "ok" {
			fmt.Println("【admin】权限验证 中间件: 验证通过")
			c.Next()
			return
		}
		fmt.Println("【admin】权限验证 中间件: 验证失败")
		//
		c.JSON(500, rg.Error[string]("权限验证失败"))
		c.Abort()
		return
	}
}
