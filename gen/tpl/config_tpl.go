package tpl

var ConfigTpl = `
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const (
	env_dev    = "dev"
	env_beta   = "beta"
	env_online = "online"
)

var (
	Viper *viper.Viper

	Env = os.Getenv("ENVIRON")
)

func Init() error {
	if Env == "" || (Env != env_beta && Env != env_online) {
		Env = env_dev
	}
	Viper = viper.New()
	Viper.AddConfigPath("./etc/app/")
	Viper.SetConfigType("toml")
	Viper.SetConfigName(Env)

	if err := Viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Sprintf("load config error: %s", err.Error()))
		return err
	}
	return nil
}
`
