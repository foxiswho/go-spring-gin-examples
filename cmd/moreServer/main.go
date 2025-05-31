package main

import (
	"github.com/go-spring/spring-core/gs"
	_ "github/foxiswho/go-spring-gin-examples/middleware/auth"
	_ "github/foxiswho/go-spring-gin-examples/middleware/db/dbPostgresql"
	_ "github/foxiswho/go-spring-gin-examples/middleware/runner"
	_ "github/foxiswho/go-spring-gin-examples/middleware/server/ginServer"
	_ "github/foxiswho/go-spring-gin-examples/middleware/server/httpServer"
	_ "github/foxiswho/go-spring-gin-examples/router"
	"os"
)

func init() {
	//指定环境
	//gs.SetActiveProfiles("dev")
	//关闭 案例 server
	gs.EnableSimplePProfServer(false)
	// 指定配置文件目录, 如果不设置，默认 conf 目录
	_ = os.Setenv("GS_SPRING_APP_CONFIG-LOCAL_DIR", "./config")
}
func main() {
	// 启动 Go-Spring 应用（自动启动 Gin 服务）
	gs.Run()
}
