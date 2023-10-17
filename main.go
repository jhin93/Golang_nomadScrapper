package main

import (
	"fmt"
	"strings"

	"github.com/jhin93/Golang/scrapper"
	"github.com/labstack/echo"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	fmt.Println(c.FormValue("term")) // html의 input에 입력된 값 체크(name="term")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323")) // http://localhost:1323/ 에서 확인 가능
}
