package service

import (
	"log"
	"net/http"
)

const (
	carServiceUrl = "https://myfakeapi.com/api/cars/1"
)

type CarService interface {
	FetchData()
}

type fetchCarDataService struct {
}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData() {
	client := http.Client{}
	log.Printf("Fetching data from %s", carServiceUrl)
	resp, _ := client.Get(carServiceUrl)
	carDatachannel <- resp
}
