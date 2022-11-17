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
	fmt.Println("проверка e-mail данных")
	service.Email()
	fmt.Println("проверка биллинга")
	service.Billing()
	fmt.Println("проверка MMS")
	service.MmsData()

}
