package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/show", show)
	e.Logger.Fatal(e.Start(":8000"))
}

func show(c echo.Context) error {
	q := c.QueryParam("q")
	return c.String(200, q)
}
