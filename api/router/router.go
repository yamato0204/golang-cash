package router

import (
	"cash/sample-cash/api/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(c controller.Controller) *echo.Echo {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.GET("/user/:id", c.GetUser)

	return e
}
