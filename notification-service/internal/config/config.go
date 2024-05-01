package config

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type App struct {
	Environment string
	EmailBackend string
	FromEmail string
	FCMServerKey string
	PushNotificationBackend string
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
	Port 	string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
	DB int
}


var RedisSetting = &Redis{}

type RabbitMQ struct {
	Username string
	Password string
	Host string
	Port string
}

var RabbitMQSettings = &RabbitMQ{}

type SMTPServer struct {
	From     string
	PassWord string
	Host     string
	Port     string
}

var SmtpSetting = &SMTPServer{}

var cfg *ini.File

// Setup initialize the configuration instance
func Setup() {
	var err error
	cfg, err = ini.Load("app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("smtp", SmtpSetting)
	mapTo("redis", RedisSetting)
	mapTo("rabbitmq", RabbitMQSettings)

	// AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	// ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	// ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	// RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
