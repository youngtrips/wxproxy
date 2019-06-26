package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type HttpInfo struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

type WXInfo struct {
	AppId     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
	Token     string `yaml:"token"`
}

type ServerInfo struct {
	Http HttpInfo `yaml:"http"`
	WX   WXInfo   `yaml:"wx"`
}

//////////////
var (
	PWD  string
	_cfg ServerInfo
)

func init() {
	PWD, _ = os.Getwd()
}

func Load(cfgfile string) error {
	data, err := ioutil.ReadFile(cfgfile)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal([]byte(data), &_cfg); err != nil {
		return err
	}
	return nil
}

func Config() ServerInfo {
	return _cfg
}
