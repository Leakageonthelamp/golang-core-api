package middlewares

import "github.com/labstack/echo/v4"

func RewritePath(path string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetPath(path)
			return next(c)
		}
	}
}
