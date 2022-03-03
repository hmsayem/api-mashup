package main

import (
	"github.com/hmsayem/go-api-mashup/controller"
	"github.com/hmsayem/go-api-mashup/router"
	"github.com/hmsayem/go-api-mashup/service"
	"log"
	"net/http"
)

var (
	mashupService    = service.NewMashupService(5)
	mashupController = controller.NewMashupController(mashupService)
	muxRouter        = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"
	muxRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server is up and running.")
	})
	muxRouter.GET("/cars", mashupController.GetMashup)
	muxRouter.SERVE(port)
}
