package service

import (
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/syslog"
	"github/foxiswho/go-spring-gin-examples/app/domainDemo/model/modUser"
	"github/foxiswho/go-spring-gin-examples/infrastructure/entityDemo"
	"github/foxiswho/go-spring-gin-examples/infrastructure/repositoryDemo"
	"github/foxiswho/go-spring-gin-examples/pkg/tools/wrapperPg/rg"
	"reflect"
)

func init() {
	gs.Provide(new(UserDbService)).Init(func(s *UserDbService) {
		syslog.Debugf("%+v initialized successfully", reflect.TypeOf(s).String())
	})
}

// 用户 服务类
type UserDbService struct {
	rep *repositoryDemo.UserRepository `autowire:""`
}

// 保存到数据库
func (t *UserDbService) Save(c *gin.Context, ct modUser.CreateUpdateCt) (rt rg.Rs[string]) {
	syslog.Infof("ct: %+v", ct)
	if ct.Name == "" {
		return rt.ErrorMessage("名称不能为空")
	}
	info := &entityDemo.UserEntity{
		Name: ct.Name,
	}
	t.rep.Save(info)
	return rt.Ok()
}
