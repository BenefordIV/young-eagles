package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	DB_NAME          = "DB_NAME"
	DB_HOST          = "DB_HOST"
	DB_PORT          = "DB_PORT"
	DB_PASS          = "DB_PASS"
	APP_ENV_LOCATION = "APP_ENV_LOCATION"
)

func LoadConfig(path string) {
	log.Println("young-eagles-config path")
	viper.AddConfigPath(path)
	viper.SetConfigName("application")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("oh no cannot read viper")
	}
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

func GetDbPass() string {
	return viper.GetString(DB_PASS)
}

func GetAppEnvLocation() string {
	loc := os.Getenv(APP_ENV_LOCATION)
	if len(loc) == 0 {
		loc = "./"
	}

	return loc
}
