package main

import (
	"net/http"
	"os"

	"github.com/diegohce/logger"
	"github.com/julienschmidt/httprouter"
)

var (
	log *logger.Logger
)

func main() {

	bindAddr := os.Getenv("BADSERVICE_BINDADDR")
	if bindAddr == "" {
		bindAddr = ":6666"
	}

	log = logger.New("badservice::")

	router := httprouter.New()
	router.GET("/badservice/drop", dropConnectionHandler)
	router.GET("/badservice/status/:code", statusCodeHandler)
	router.GET("/badservice/delay/:delay", delayHandler)
	router.GET("/badservice/headers", showHeadersHandler)

	router.POST("/badservice/drop", dropConnectionHandler)
	router.POST("/badservice/status/:code", statusCodeHandler)
	router.POST("/badservice/delay/:delay", delayHandler)
	router.POST("/badservice/headers", showHeadersHandler)

	router.PUT("/badservice/drop", dropConnectionHandler)
	router.PUT("/badservice/status/:code", statusCodeHandler)
	router.PUT("/badservice/delay/:delay", delayHandler)
	router.PUT("/badservice/headers", showHeadersHandler)

	router.DELETE("/badservice/drop", dropConnectionHandler)
	router.DELETE("/badservice/status/:code", statusCodeHandler)
	router.DELETE("/badservice/delay/:delay", delayHandler)
	router.DELETE("/badservice/headers", showHeadersHandler)

	log.Info().Println("Starting badservice on", bindAddr)

	log.Error().Fatal(http.ListenAndServe(bindAddr, router))
}
