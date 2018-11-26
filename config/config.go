package config

import "github.com/jinzhu/configor"

// Config Payment Gateway
var Config = struct {
	App struct {
		ENV      string
		HTTPAddr string
		HTTPPort string
	}

	DB struct {
		Driver   string
		Host     string
		Port     string
		Name     string
		User     string
		Password string
		Locale   string
	}

	Redis struct {
		Host     string
		Port     string
		Password string
		DB       int
		Duration int // in minute
	}
}{}

func init() {
	configor.Load(&Config, "config.yaml")
}
