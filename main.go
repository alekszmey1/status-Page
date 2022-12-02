package main

import (
	"awesomeProject/skillbox/StatusPage/service"
	"log"
)

func main() {

	log.Println("запуск программы")
	/*log.Println("проверка смс данных")
	service.SmsData() /*
		log.Println("проверка голосовых данных")
		service.VoiceCall()
		log.Println("проверка e-mail данных")
		service.Email()
		log.Println("проверка биллинга")
		service.Billing()
		log.Println("проверка MMS")
		service.MmsData()
		log.Println("проверка support")
		service.Support()
		log.Println("проверка incident")
		service.Incident()*/
	log.Println("запуск сортировки смс")
	service.SortSMSOne()
}
