package main

import (
	"strconv"
	"url_shortener/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	port = 8080
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/shorten", controller.PostShortenController)
	e.GET("/:shortCode", controller.GetShortenCodeController)
	e.GET("/stats/:shortCode", controller.GetShortenCodeStatsController)
	e.DELETE("/shorten/:shortCode", controller.DeleteShortenCodeController)

	var address string = ":" + strconv.Itoa(port)
	e.Logger.Fatal(e.Start(address))
}
