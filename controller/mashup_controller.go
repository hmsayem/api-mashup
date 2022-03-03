package controller

import (
	"encoding/json"
	"github.com/hmsayem/go-api-mashup/service"
	"log"
	"net/http"
)

var (
	mashupService service.MashupService
)

type MashupController interface {
	GetMashup(response http.ResponseWriter, request *http.Request)
}
type controller struct{}

func NewMashupController(service service.MashupService) MashupController {
	mashupService = service
	return &controller{}
}
func (*controller) GetMashup(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	result, err := mashupService.GetList()
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
