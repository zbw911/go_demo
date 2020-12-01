package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {

}

// ...
const (
	kAppName       = "APP_NAME"
	kConfigServer  = "CONFIG_SERVER"
	kConfigLabel   = "CONFIG_LABEL"
	kConfigProfile = "CONFIG_PROFILE"
	kConfigType    = "CONFIG_TYPE"
)

func loadRemoteConfig() (err error) {
	// 组装配置文件地址: http://localhost:8080/config/go-app-default.yaml
	confAddr := fmt.Sprintf("%v/%v/%v-%v.yml",
		viper.Get(kConfigServer), viper.Get(kConfigLabel),
		viper.Get(kAppName), viper.Get(kConfigProfile))
	resp, err := http.Get(confAddr)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 设置配置文件格式: yaml
	viper.SetConfigType(viper.GetString(kConfigType))
	// 载入配置文件
	if err = viper.ReadConfig(resp.Body); err != nil {
		return
	}
	log.Println("Load config from: ", confAddr)
	return
}
