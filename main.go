package main

import (
	"fmt"
	"log"

	"main/config"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DBName)
	fmt.Println(config.Config.LogFile)
	log.Println("test")
}
