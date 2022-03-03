package service

import (
	"encoding/json"
	"github.com/hmsayem/go-api-mashup/entity"
	"net/http"
)

type CarDetailsService interface {
	GetDetails() (*entity.CarDetails, error)
}

var (
	carService       = NewCarService()
	ownerService     = NewOwnerService()
	carDatachannel   = make(chan *http.Response)
	ownerDatachannel = make(chan *http.Response)
)

type service struct {
}

func NewCarDetailsService() CarDetailsService {
	return &service{}
}

func (*service) GetDetails() (*entity.CarDetails, error) {

	go carService.FetchData()
	go ownerService.FetchData()
	var carData, ownerData *http.Response

	for i := 0; i < 2; i++ {
		select {
		case carResponse := <-carDatachannel:
			carData = carResponse
		case ownerResponse := <-ownerDatachannel:
			ownerData = ownerResponse
		}
	}
	var car entity.Car
	if err := json.NewDecoder(carData.Body).Decode(&car); err != nil {
		return nil, err
	}
	var owner entity.Owner
	if err := json.NewDecoder(ownerData.Body).Decode(&owner); err != nil {
		return nil, err
	}

	return &entity.CarDetails{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}, nil
}