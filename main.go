package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.Logger.Fatal(e.Start(":1323")) // http://localhost:1323/ 에서 확인 가능
}
