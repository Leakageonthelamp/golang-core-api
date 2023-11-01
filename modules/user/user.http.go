package user

import (
	core "github.com/Leakageonthelamp/go-leakage-core"
	"github.com/labstack/echo/v4"
)

func NewUserHTTP(e *echo.Group) {
	controller := &UserController{}

	e.POST("/user", core.WithHTTPContext(controller.Create))
	e.GET("/user", core.WithHTTPContext(controller.Pagination))
	e.GET("/user/:id", core.WithHTTPContext(controller.Find))
	e.PUT("/user/:id", core.WithHTTPContext(controller.Update))
	e.DELETE("/user/:id", core.WithHTTPContext(controller.Delete))
}
