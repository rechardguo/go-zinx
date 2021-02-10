package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	IP             string
	Port           uint32
	Name           string
	MaxConn        uint32
	MaxPackageSize uint32
	Version        string
}

func (self *Conf) Reload() {
	data, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, Config)
}

func init() {
	Config = &Conf{
		IP:             "0.0.0.0",
		Port:           9000,
		Name:           "Zinx Server",
		MaxConn:        100,
		MaxPackageSize: 2048,
		Version:        "v0.4",
	}
	//从配置文件里去读取
	//Config.Reload()
}

var Config *Conf
