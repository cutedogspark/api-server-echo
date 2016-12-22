package config

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/cutedogspark/api-server-echo/service/Echo"
)

type (
	MariaDB struct {
		Name     string `json:"name"`
		User     string `json:"user" default:"root"`
		Password string `json:"password" required:"true" env:"db_password"`
		Host     string `json:"host" default:"localhost"`
		Port     string `json:"port" default:"3306"`
	}

	Production struct {
		MariaDB    MariaDB
	}

	Development struct {
		MariaDB    MariaDB
	}

	Test struct {
		MariaDB    MariaDB
	}
)

func Init(x *Echo.Service) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err!= nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	fmt.Println("config read success")
}


