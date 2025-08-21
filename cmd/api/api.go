package main

import (
	"plaque/engine"

	"github.com/labstack/echo/v4"
)

func main() {
	s := echo.New()

	s.GET("p", func(c echo.Context) error {

		y := c.QueryParam("y")

		w := c.QueryParam("w")

		if y == "" && w == "" {

			return c.String(echo.ErrBadRequest.Code,
				"city and province is required is query param like this: ?y=س&w=11",
			)

		}

		reciveData := engine.Engine(y, w)

		return c.JSON(200, reciveData)
	})

	s.Logger.Fatal(s.Start(":4000"))
}
