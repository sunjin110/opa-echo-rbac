package emiddleware

import (
	"net/http"
	"opa-echo-test/infrastructure/rbac"
	"opa-echo-test/internal/chk"

	"github.com/labstack/echo/v4"
)

// Authorization 認可
func Authorization(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		// JWTからユーザーを識別できるIDを取得しているということを仮定する

		// TODO ユーザーの権限をDBから取得する

		// RBACの評価inputを作成
		input := rbac.Input{
			User:              "user",
			Roles:             []string{"read-only"},
			AllowResourceList: []string{"test-app"},

			Method:         c.Request().Method, // GET,POST,PUT,DELETE
			Path:           c.Path(),
			AccessResource: "test-app", // これは、pathから頑張って取得する必要がある
		}

		// 許可されていない場合は、403
		if !rbac.Eval(c.Request().Context(), input) {
			return c.JSON(http.StatusForbidden, "許可されていません")
		}

		chk.SE(next(c))

		return nil
	}

}
