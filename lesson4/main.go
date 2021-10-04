package main

import (
	"lesson4/config"
	"lesson4/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}
