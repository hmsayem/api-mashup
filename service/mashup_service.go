package service

import (
	"encoding/json"
	"github.com/hmsayem/go-api-mashup/entity"
	"net/http"
)

type MashupService interface {
	GetList() ([]entity.Mashup, error)
}

var (
	carService       = NewCarService()
	ownerService     = NewOwnerService()
	carDatachannel   = make(chan *http.Response)
	ownerDatachannel = make(chan *http.Response)
	count            int
)

type service struct{}

func NewMashupService(cnt int) MashupService {
	count = cnt
	return &service{}
}

func (*service) GetList() ([]entity.Mashup, error) {
	var cars []entity.Mashup
	for i := 1; i <= count; i++ {
		car, err := GetMashup(i)
		if err != nil {
			return nil, err
		}
		cars = append(cars, *car)
	}
	return cars, nil
}

func GetMashup(id int) (*entity.Mashup, error) {
	go carService.FetchData(id)
	go ownerService.FetchData(id)
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

	return &entity.Mashup{
		ID:        car.ID,
		Brand:     car.Brand,
		Model:     car.Model,
		Year:      car.Year,
		FirstName: owner.FirstName,
		LastName:  owner.LastName,
		Email:     owner.Email,
	}, nil
}
