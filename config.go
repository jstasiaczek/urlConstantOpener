package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/kardianos/osext"
)

var ConfigPath string
var App AppConfig

func init() {
	execDir, _ := osext.ExecutableFolder()
	execDir += "/config.json"
	flag.StringVar(&ConfigPath, "c", execDir, "Config file location, default is exec dir + config.json")
	flag.Parse()
	App = NewConfig()
	App.ImportSettings(ConfigPath)
}

type AppConfig struct {
	DirectoryToScan string
	LogFile         string
	UseFiles        bool
	UseHttp         bool
	UseFileLog      bool
	SecretKey       string
	HttpPort        int
}

func (this *AppConfig) ImportSettings(path string) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		panic("Can't open config file! \n" + err.Error())
	}
	err = json.Unmarshal(raw, this)
	if err != nil {
		panic("Can't read config file! \n" + err.Error())
	}
}

func NewConfig() AppConfig {
	return AppConfig{}
}
