package services

import (
	"test_package/models"
	"test_package/requests"

	core "github.com/Leakageonthelamp/go-leakage-core"
	"github.com/Leakageonthelamp/go-leakage-core/errmsgs"
	core_models "github.com/Leakageonthelamp/go-leakage-core/models"
	"github.com/Leakageonthelamp/go-leakage-core/repository"
	"github.com/Leakageonthelamp/go-leakage-core/utils"
)

type IUserService interface {
	Create(request *requests.UserCreateRequest) (*models.User, core.IError)
	Paginate(pageOptions *core_models.PageOptions) (*repository.Pagination[models.User], core.IError)
	Find(ID string) (*models.User, core.IError)
	Update(ID string, request *requests.UserUpdateRequest) (*models.User, core.IError)
	Delete(ID string) core.IError
}

type userService struct {
	ctx core.IContext
}

func NewUserService(ctx core.IContext) IUserService {
	return &userService{ctx}
}

func (s *userService) Create(request *requests.UserCreateRequest) (*models.User, core.IError) {
	payload := &models.User{
		BaseModel: models.NewBaseModel(),
	}
	if err := utils.Copy(payload, request); err != nil {
		return nil, s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	ierr := repository.New[models.User](s.ctx).Create(payload)
	if ierr != nil {
		return nil, ierr
	}

	return s.Find(payload.ID)
}

func (s *userService) Paginate(pageOptions *core_models.PageOptions) (*repository.Pagination[models.User], core.IError) {
	resp, ierr := repository.New[models.User](s.ctx).Pagination(pageOptions)
	if ierr != nil {
		return nil, ierr
	}

	return resp, nil
}

func (s *userService) Find(ID string) (*models.User, core.IError) {
	resp, ierr := repository.New[models.User](s.ctx).FindOne("id = ?", ID)
	if ierr != nil {
		return nil, ierr
	}

	return resp, nil
}

func (s *userService) Update(ID string, request *requests.UserUpdateRequest) (*models.User, core.IError) {
	user, ierr := s.Find(ID)
	if ierr != nil {
		return nil, ierr
	}

	if err := utils.Copy(user, request); err != nil {
		return nil, s.ctx.NewError(err, errmsgs.InternalServerError)
	}

	ierr = repository.New[models.User](s.ctx).Where("id = ?", ID).Updates(user)
	if ierr != nil {
		return nil, ierr
	}

	return s.Find(ID)
}

func (s *userService) Delete(ID string) core.IError {
	ierr := repository.New[models.User](s.ctx).Where("id = ?", ID).Delete()
	if ierr != nil {
		return ierr
	}

	return nil
}
