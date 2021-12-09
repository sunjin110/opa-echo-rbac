package emiddleware

import (
	"net/http"
	"opa-echo-test/infrastructure/rbac"
	"opa-echo-test/infrastructure/sqlite"
	"opa-echo-test/infrastructure/sqlite/repository"
	"opa-echo-test/internal/chk"

	"github.com/labstack/echo/v4"
)

// Authorization 認可
func Authorization(next echo.HandlerFunc) echo.HandlerFunc {

	repo := repository.NewUserAuthInfoRepository(sqlite.GetDB())

	return func(c echo.Context) error {

		// JWTからユーザーを識別できるIDとaccessResourceを取得していると仮定する
		userID := uint64(1)
		accessResource := "test-app-1"

		// ユーザーの権限をDBから取得する
		userAuthInfo := repo.Get(userID)

		// RBACの評価inputを作成
		input := rbac.Input{
			User:              userAuthInfo.Name,
			Roles:             userAuthInfo.RoleList,
			AllowResourceList: userAuthInfo.ResourceList,

			Method:         c.Request().Method, // GET,POST,PUT,DELETE
			Path:           c.Path(),
			AccessResource: accessResource,
		}

		// 許可されていない場合は、403
		if !rbac.Eval(c.Request().Context(), input) {
			return c.JSON(http.StatusForbidden, "許可されていません")
		}

		chk.SE(next(c))

		return nil
	}

}
