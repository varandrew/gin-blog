package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Conf *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Conf, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Conf.Section("server").Key("RunMode").MustString("debug")
}

func LoadServer() {
	sec, err := Conf.GetSection("server")
	if err != nil {
		log.Fatal("Fail to get section 'server' : %v", err)
	}

	HTTPPort = sec.Key("HttpPort").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("ReadTimeout").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WriteTimeout").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Conf.GetSection("app")
	if err != nil {
		log.Fatal("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JwtSecret").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PageSize").MustInt(10)
}
