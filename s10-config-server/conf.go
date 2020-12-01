package main

import (
	viper "github.com/spf13/viper"
	"log"
	"os"
)

func init() {
	viper.SetConfigName("go-app")
	viper.AddConfigPath(os.ExpandEnv(`$GOPATH\src\github.com\GotaX\config-server-demo`))
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Fail to load config", err)
	}
}
