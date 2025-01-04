package main

import (
	"fmt"
	"young-eagles/internal/config"
)

func main() {
	config.LoadConfig(config.GetAppEnvLocation())

	fmt.Println(config.GetDbPort())
}
