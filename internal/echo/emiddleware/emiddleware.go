package emiddleware

import (
	"log"
	"opa-echo-test/internal/chk"
	"opa-echo-test/internal/jsonutil"

	"github.com/labstack/echo/v4"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		// Authorization

		log.Println("path is ", jsonutil.Marshal(c.Request().URL.Path))

		err := next(c)
		chk.SE(err)

		return nil
	}

}
