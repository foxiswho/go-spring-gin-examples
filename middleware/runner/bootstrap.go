package runner

import (
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/syslog"
	"github/foxiswho/go-spring-gin-examples/pkg/config"
)

func init() {
	gs.Object(&Bootstrap{}).AsRunner()
}

type Bootstrap struct {
	ser config.Server `value:"${server}"`
}

func (b *Bootstrap) Run() error {
	syslog.Infof("starting successfully. Port: %+v", b.ser.Port)
	return nil
}
