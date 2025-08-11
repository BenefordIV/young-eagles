package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	DB_NAME = "DB_NAME"
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_PASS = "DB_PASS"
	DB_USER = "DB_USER"
)

func LoadConfig() {
	log.Println("young-eagles-config path")
	viper.AutomaticEnv()
	log.Println("config loaded")
}

func GetDbName() string {
	return viper.GetString(DB_NAME)
}

func GetDbHost() string {
	return viper.GetString(DB_HOST)
}

func GetDbPort() int {
	return viper.GetInt(DB_PORT)
}

func GetDbUsername() string { return viper.GetString(DB_USER) }

func GetDbPass() string {
	return viper.GetString(DB_PASS)
}
