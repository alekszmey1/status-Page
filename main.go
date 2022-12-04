package main

import (
	"awesomeProject/skillbox/StatusPage/server"
	"log"
)

func main() {

	log.Println("запуск программы")
	server.App()
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
	/*log.Println("запуск сортировки смс")
	service.SortSMSOne()
	log.Println("запуск сортировки mмс")
	service.SortMMSOne()
	log.Println("запуск сортировки  email ")
	service.SortEmail()
	log.Println("запуск сортировки support ")
	service.SortSupport()
	log.Println("запуск сортировки support ")
	service.SortIncident()*/
}
