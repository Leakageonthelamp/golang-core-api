package user

import (
	"net/http"
	"test_package/requests"
	"test_package/services"

	core "github.com/Leakageonthelamp/go-leakage-core"
)

type UserController struct{}

func (u *UserController) Create(c core.IHTTPContext) error {
	request := &requests.UserCreateRequest{}
	if err := c.BindWithValidate(request); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	userSvc := services.NewUserService(c)
	response, err := userSvc.Create(request)
	if err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *UserController) Pagination(c core.IHTTPContext) error {
	userSvc := services.NewUserService(c)
	response, err := userSvc.Paginate(c.GetPageOptions())
	if err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) Find(c core.IHTTPContext) error {
	id := c.Param("id")

	userSvc := services.NewUserService(c)
	response, err := userSvc.Find(id)
	if err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) Update(c core.IHTTPContext) error {
	request := &requests.UserUpdateRequest{}
	if err := c.BindWithValidate(request); err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	userSvc := services.NewUserService(c)
	response, err := userSvc.Update(c.Param("id"), request)
	if err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserController) Delete(c core.IHTTPContext) error {
	id := c.Param("id")

	userSvc := services.NewUserService(c)
	err := userSvc.Delete(id)
	if err != nil {
		return c.JSON(err.GetStatus(), err.JSON())
	}

	return c.JSON(http.StatusAccepted, nil)
}
