package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	PrefixUrl     string
	JwtSecret     string
	JwtExpireMins time.Duration
	AppName       string
}

var AppSetting = &App{}

type Database struct {
	DatabaseUri  string
	DatabaseName string
}

var DbSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)

	mapTo("database", DbSetting)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
