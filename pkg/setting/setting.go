package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type appSetting struct {
	PrefixUrl           string
	JwtSecret           string
	JwtExpireMins       time.Duration
	AppName             string
	ItemPageSize        int
	ListingPageSize     int
	TransactionPageSize int
}

var App = &appSetting{}

type serverSetting struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var Server = &serverSetting{}

type databaseSetting struct {
	DatabaseUri  string
	DatabaseName string
}

var DB = &databaseSetting{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", App)

	mapTo("server", Server)

	mapTo("database", DB)

	// post format
	App.JwtExpireMins *= time.Minute
	Server.ReadTimeout *= time.Second
	Server.WriteTimeout *= time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
