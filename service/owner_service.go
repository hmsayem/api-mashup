package service

import (
	"log"
	"net/http"
)

const (
	ownerServiceUrl = "https://myfakeapi.com/api/users/1"
)

type OnwerService interface {
	FetchData()
}

type fetchOwnerDataService struct {
}

func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData() {
	client := http.Client{}
	log.Printf("Fetching data from %s", ownerServiceUrl)
	resp, _ := client.Get(ownerServiceUrl)
	ownerDatachannel <- resp
}
