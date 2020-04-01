package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
	Conf, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = Conf.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Conf.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize *= 1024 * 1024

	err = Conf.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Conf.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	err = Conf.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Conf.MapTo DatabaseSetting err: %v", err)
	}
}
