package main

import (
	_ "embed"
	"fmt"
	"opa-echo-test/controller"
	"opa-echo-test/domain/entity"
	"opa-echo-test/infrastructure/rbac"
	"opa-echo-test/infrastructure/sqlite"
	"opa-echo-test/infrastructure/sqlite/repository"
	"opa-echo-test/internal/chk"
	"opa-echo-test/internal/echo/emiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	dbPath = "./db/test.db"
	dbDir  = "./db"
)

//go:embed policy/rbac.rego
var opaRbacModule []byte

func main() {
	fmt.Println("opaを使用して、権限のアクセスがうまく動いていることを確認してみる、検証")

	// sqlite setup
	sqlite.Setup(dbDir, dbPath)

	// setup user
	// テストするために、ユーザーデータを作成する
	setupTestData()

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
	e.GET("/apps/:id", controller.AppDetailGet)
	e.POST("/apps", controller.AppsPost)

	err := e.Start(":1234")
	chk.SE(err)
}

// テスト用のデータを作る
func setupTestData() {

	repo := repository.NewUserAuthInfoRepository(sqlite.GetDB())

	// read-only
	readOnlyUser := &entity.UserAuthInfo{
		UserID:       1,
		Name:         "read-only-user",
		RoleList:     []string{"read-only"},
		ResourceList: []string{"test-app-1"},
	}

	adminUser := &entity.UserAuthInfo{
		UserID:       2,
		Name:         "admin-user",
		RoleList:     []string{"admin"},
		ResourceList: []string{".*"},
	}

	repo.Insert(readOnlyUser)
	repo.Insert(adminUser)
}
