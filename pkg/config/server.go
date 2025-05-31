package config

// Server 服务配置
type Server struct {
	Port   int    `value:"${port:=18080}"`
	Domain string `value:"${domain:=}"`
}
