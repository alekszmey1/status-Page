package main

import (
	"awesomeProject/skillbox/StatusPage/server"

	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	//var WG sync.WaitGroup
	//WG.Add(1)
	log.Info("запуск программы")

	server.App()
	log.Info("отработал сервер")
	//WG.Done()
	//WG.Wait()
}
