package middlewares

import (
	"strings"
	"test_package/consts"
	"test_package/services"

	core "github.com/Leakageonthelamp/go-leakage-core"
	"github.com/Leakageonthelamp/go-leakage-core/errmsgs"
	"github.com/labstack/echo/v4"
	"github.com/pskclub/mine-core/utils"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(core.IHTTPContext)

		authorization := strings.TrimSpace(cc.Request().Header.Get("Authorization"))
		if authorization == "" {
			return c.JSON(401, "Authorization header is required")
		}

		userSvc := services.NewUserService(cc)
		user, ierr := userSvc.Find(authorization)
		if errmsgs.IsNotFoundError(ierr) {
			return c.JSON(401, "Token is invalid")
		}
		if ierr != nil {
			return c.JSON(ierr.GetStatus(), ierr.JSON())
		}

		c.Set(consts.ContextKeyUser, user)
		c.Set(consts.ContextKeyUserToken, authorization)

		cc.SetUser(&core.ContextUser{
			ID:   user.ID,
			Name: utils.GetString(&user.FirstName) + " " + utils.GetString(&user.LastName),
		})

		return next(c)
	}
}
