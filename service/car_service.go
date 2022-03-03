package service

import (
	"log"
	"net/http"
	"strconv"
)

var (
	carServiceUrl = "https://myfakeapi.com/api/cars/"
)

type CarService interface {
	FetchData(id int)
}

type fetchCarDataService struct{}

func NewCarService() CarService {
	return &fetchCarDataService{}
}

func (*fetchCarDataService) FetchData(id int) {
	client := http.Client{}
	log.Printf("Fetching data from %s", carServiceUrl+strconv.Itoa(id))
	resp, _ := client.Get(carServiceUrl + strconv.Itoa(id))
	carDatachannel <- resp
}
