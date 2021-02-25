package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/hello", hello)
	e.GET("/sendData", sendData)
	e.GET("/sendData2", sendData2)
	e.GET("/getData", getData)
	e.Logger.Fatal(e.Start(":8090"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func sendData(c echo.Context) error {
	index := c.FormValue("index")
	data := c.FormValue("data")
	fmt.Println("sendData -- " + index + ":" + data)
	send(index, data)
	return c.String(http.StatusOK, "Hello, World!")
}

func sendData2(c echo.Context) error {
	index := c.FormValue("index")
	data := c.FormValue("data")
	fmt.Println("sendData2 -- " + index + ":" + data)
	send2(index, data)
	return c.String(http.StatusOK, "Hello, World!")
}

func getData(c echo.Context) error {
	index := c.FormValue("index")
	data := get(index)
	fmt.Println("getData -- " + index + ":" + data)
	return c.String(http.StatusOK, data)
}
