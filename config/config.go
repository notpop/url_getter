package config

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

type ConfigList struct {
	TargetUrl            string
	GetHtmlPath          string
	SearchLimit          int
	OriginSourceSelector string
	SubSelector          string
	DbName               string
	SQLDriver            string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("faild to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		TargetUrl:            cfg.Section("web").Key("target_url").String(),
		GetHtmlPath:          cfg.Section("web").Key("get_html_path").String(),
		SearchLimit:          cfg.Section("web").Key("search_limit").MustInt(),
		OriginSourceSelector: cfg.Section("web").Key("origin_source_selector").String(),
		SubSelector:          cfg.Section("web").Key("sub_selector").String(),
		DbName:               cfg.Section("db").Key("name").String(),
		SQLDriver:            cfg.Section("db").Key("driver").String(),
	}
}
