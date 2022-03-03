package controller

import (
	"encoding/json"
	"github.com/hmsayem/go-api-mashup/service"
	"log"
	"net/http"
)

var (
	carDetailService service.CarDetailsService
)

type CarDetailsController interface {
	GetCarDetails(response http.ResponseWriter, request *http.Request)
}
type controller struct{}

func NewCarDetailsController(service service.CarDetailsService) CarDetailsController {
	carDetailService = service
	return &controller{}
}
func (*controller) GetCarDetails(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	result, err := carDetailService.GetDetails()
	if err != nil {
		log.Printf("failed to get car details: %v", err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(result); err != nil {
		return
	}
}
