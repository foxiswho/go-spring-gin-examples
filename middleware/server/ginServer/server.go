package ginServer

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util/syslog"
	"net"
	"net/http"
)

// GinServerDefault 初始化默认服务
var GinServerDefault = gin.New()

func init() {
	//服务，传入配置 内端口
	gs.Provide(NewGinServer, gs.TagArg("${server.port}")).AsServer()
	//静态文件夹
	GinServerDefault.Static("/assets", "./assets")
}

// gin 框架 整合
type GinServer struct {
	svr       *http.Server
	svrEngine *gin.Engine
}

func NewGinServer(port string) *GinServer {
	syslog.Debugf("NewGinServer.port:%+v ", port)
	svr := &GinServer{}
	svr.svrEngine = GinServerDefault
	svr.svr = &http.Server{
		Addr:    ":" + port,
		Handler: svr.svrEngine,
	}
	return svr
}

// 启动 端口
func (s *GinServer) ListenAndServe(sig gs.ReadySignal) error {
	addr := s.svr.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	<-sig.TriggerAndWait() // 等待启动信号
	//
	syslog.Infof("starting successfully... Port: %+v", addr)
	return s.svr.Serve(ln)
}

// 关闭
func (s *GinServer) Shutdown(ctx context.Context) error {
	return s.svr.Shutdown(ctx)
}
