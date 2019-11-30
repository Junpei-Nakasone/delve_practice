package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from echo.")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	return c.String(http.StatusOK, fmt.Sprintf("your cat name is: %s and %s", catName, catType))
}

func main() {
	fmt.Println("hello world from web server.")

	e := echo.New()

	e.GET("/", hello)
	e.GET("/cats", getCats)
	e.Start(":8080")
}
