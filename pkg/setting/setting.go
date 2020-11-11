package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	// Cfg 配置文件
	Cfg *ini.File

	// RunMode 运行模式
	RunMode string

	// HTTPPort 运行端口
	HTTPPort int
	// ReadTimeout 读取超时
	ReadTimeout time.Duration
	// WriteTimeout 写入超时
	WriteTimeout time.Duration

	// PageSize 页面大小
	PageSize int
	// JwtSecret token口令
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")

	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase 加载基础配置
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 加载服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")

	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp 加载应用配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")

	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("mengyuxu")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
