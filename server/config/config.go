package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var config Config

type Config struct {
	Db []string
}

func init() {
	bytes,err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic("读取配置文件异常："+err.Error())
	}


	config := &Config{}
	if err := yaml.Unmarshal(bytes,config);err != nil{
		panic("反序列化配置文件异常："+err.Error())
	}
}

func GetConfig() Config {
	return config
}
