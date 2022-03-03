package service

import (
	"log"
	"net/http"
	"strconv"
)

var (
	ownerServiceUrl = "https://myfakeapi.com/api/users/"
)

type OwnerService interface {
	FetchData(id int)
}

type fetchOwnerDataService struct {
}

func NewOwnerService() CarService {
	return &fetchOwnerDataService{}
}

func (*fetchOwnerDataService) FetchData(id int) {
	client := http.Client{}
	log.Printf("Fetching data from %s", ownerServiceUrl+strconv.Itoa(id))
	resp, _ := client.Get(ownerServiceUrl + strconv.Itoa(id))
	ownerDatachannel <- resp
}
