package controller

import (
	"github.com/gin-gonic/gin"
	"github/foxiswho/go-spring-gin-examples/pkg/config"
	"github/foxiswho/go-spring-gin-examples/pkg/tools/wrapperPg/rg"
	"net/http"
)

func init() {

}

// 首页
type Index struct {
	ser config.Server `value:"${server}"`
}

// 首页
func (t *Index) Index(c *gin.Context) {
	data := gin.H{
		"tpl":    "Index",
		"port":   t.ser.Port,
		"Domain": t.ser.Domain,
	}
	c.JSON(http.StatusOK, rg.OkData(data))
}

// 端口信息
func (t *Index) Port(c *gin.Context) {
	data := gin.H{
		"port":   t.ser.Port,
		"Domain": t.ser.Domain,
	}
	c.JSON(http.StatusOK, rg.OkData(data))
}
