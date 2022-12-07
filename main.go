package main

import (
	"awesomeProject/skillbox/StatusPage/server"

	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	log.Info("запуск программы")
	server.App()
}
