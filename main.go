package main

import (
	_ "embed"
	"fmt"
	"opa-echo-test/controller"
	"opa-echo-test/infrastructure/rbac"
	"opa-echo-test/internal/chk"
	"opa-echo-test/internal/echo/emiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed policy/rbac.rego
var opaRbacModule []byte

func main() {
	fmt.Println("opaを使用して、権限のアクセスがうまく動いていることを確認してみる、検証")

	// rbac setup
	rbac.Setup(opaRbacModule)

	serve()
}

func serve() {
	// echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(emiddleware.Authorization)

	e.GET("", controller.IndexGet)
	e.GET("/apps", controller.AppsGet)
	e.POST("/apps", controller.AppsPost)

	err := e.Start(":1234")
	chk.SE(err)
}
