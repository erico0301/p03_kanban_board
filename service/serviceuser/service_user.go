package serviceuser

import (
	"p03_kanban_board/constant"
	"p03_kanban_board/entity"
	"p03_kanban_board/helper"
	"p03_kanban_board/model/modeluser"
	"p03_kanban_board/repository/repositoryuser"
	"p03_kanban_board/validation"

	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type ServiceUser interface {
	Create(data modeluser.Request) (modeluser.Response, error)
	Login(data modeluser.RequestLogin) (modeluser.ResponseLogin, error)
	Update(data modeluser.Request) (modeluser.Response, error)
	DeleteByID(id uint) error
}

type service struct {
	repo repositoryuser.RepositoryUser
}

func New(repo repositoryuser.RepositoryUser) ServiceUser {
	return &service{repo: repo}
}

func (s *service) Create(data modeluser.Request) (modeluser.Response, error) {
	err := validation.ValidateUserCreate(data, s.repo)
	if err != nil {
		return modeluser.Response{}, err
	}

	entityUser := new(entity.User)

	copier.Copy(&entityUser, &data)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(entityUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return modeluser.Response{}, err
	}
	entityUser.Password = string(hashedPassword)
	entityUser.Role = constant.MemberRole

	createdUser, err := s.repo.Create(*entityUser)
	if err != nil {
		return modeluser.Response{}, err
	}

	resp := modeluser.Response{}

	copier.Copy(&resp, &createdUser)
	resp.UpdatedAt = nil

	return resp, nil
}

func (s *service) Login(data modeluser.RequestLogin) (modeluser.ResponseLogin, error) {
	err := validation.ValidateUserLogin(data)
	if err != nil {
		return modeluser.ResponseLogin{}, err
	}

	dataUser, err := s.repo.Login(data.Email)
	if err != nil {
		return modeluser.ResponseLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(data.Password))
	if err != nil {
		return modeluser.ResponseLogin{}, constant.ErrorInvalidLogin
	}

	token, err := helper.NewJwt(dataUser.ID)
	if err != nil {
		return modeluser.ResponseLogin{}, err
	}

	resp := modeluser.ResponseLogin{}
	resp.Token = token

	return resp, nil
}

func (s *service) Update(data modeluser.Request) (modeluser.Response, error) {
	err := validation.ValidateUserUpdate(data)
	if err != nil {
		return modeluser.Response{}, err
	}

	entityUser := entity.User{}
	copier.Copy(&entityUser, &data)

	updatedUser, err := s.repo.Update(entityUser)
	if err != nil {
		return modeluser.Response{}, err
	}

	resp := modeluser.Response{}

	copier.Copy(&resp, &updatedUser)

	return resp, nil
}

func (s *service) DeleteByID(id uint) error {
	return s.repo.DeleteByID(id)
}
