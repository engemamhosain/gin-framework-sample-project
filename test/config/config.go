// config/config.go
package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.SetConfigName("config")
	Config.AddConfigPath("./config")
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
	}

	env := Config.GetString("app.environment")
	if envConfigFile := "config." + env + ".yaml"; env != "" {
		Config.SetConfigName(envConfigFile)
		if err := Config.MergeInConfig(); err != nil {
			log.Printf("Error reading environment-specific config file: %s", err)
		}
	}
}
