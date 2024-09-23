package controller

import (
	"cash/sample-cash/api/usecase"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetUser(c echo.Context) error
}

type controller struct {
	u usecase.Usecase
}

func NewController(c usecase.Usecase) Controller {
	return &controller{c}
}

func (c *controller) GetUser(ctx echo.Context) error {
	userID := ctx.Param("id")

	user, err := c.u.GetUser(ctx.Request().Context(), userID)
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(200, user)

}
