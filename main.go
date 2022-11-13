package main

import (
	"awesomeProject/skillbox/StatusPage/service"
	"fmt"
)

func main() {
	fmt.Println("запуск программы")
	fmt.Println("проверка смс данных")
	service.SmsData()
	fmt.Println("проверка голосовых данных")
	service.VoiceCall()
}
