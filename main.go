package main

import (
	"fmt"
	"net/http"
	"opa-echo-test/internal/chk"
	"opa-echo-test/internal/echo/emiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("opaを使用して、権限のアクセスがうまく動いていることを確認してみる、検証")

	// echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(emiddleware.Authorization)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	err := e.Start(":1234")
	chk.SE(err)

}
