package main

import (
	"fmt"
	"log"
	"young-eagles/internal/config"
	"young-eagles/internal/services"
)

func main() {
	config.LoadConfig(config.GetAppEnvLocation())

	dbConfig := services.DbConfig{
		DbName: config.GetDbName(),
		DbUser: config.GetDbUsername(),
		DbPass: config.GetDbPass(),
		DbPort: config.GetDbPort(),
		DbHost: config.GetDbHost(),
	}

	dbConn, err := services.InitDatabase(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.PingDb()
	if err != nil {
		log.Fatalf("error pinging error %v", err)
	}
	fmt.Printf("successfully connected to %s", dbConfig.DbName)
}
