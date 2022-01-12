package servicetask

import (
	"p03_kanban_board/entity"
	"p03_kanban_board/model/modeltask"
	"p03_kanban_board/repository/repositorytask"
	"p03_kanban_board/validation"

	"github.com/jinzhu/copier"
)

type ServiceTask interface {
	Create(request modeltask.Request) (modeltask.ResponseStore, error)
	Gets() ([]modeltask.ResponseGet, error)
	Update(request modeltask.RequestUpdate) (modeltask.ResponseStore, error)
	UpdateStatus(request modeltask.RequestUpdateStatus) (modeltask.ResponseStore, error)
	UpdateCategory(request modeltask.RequestUpdateCategory) (modeltask.ResponseStore, error)
	Delete(id uint, userID uint) error
}

type Service struct {
	repo repositorytask.RepositoryTask
}

func (s *Service) Delete(id uint, userID uint) error {
	entityTask := entity.Task{
		ID:     id,
		UserID: userID,
	}
	err := s.repo.Delete(entityTask)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateStatus(request modeltask.RequestUpdateStatus) (modeltask.ResponseStore, error) {
	entityTask := entity.Task{}
	copier.Copy(&entityTask, &request)
	update, err := s.repo.Update(entityTask)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}
	resp := modeltask.ResponseStore{}
	copier.Copy(&resp, &update)
	return resp, nil
}

func (s *Service) UpdateCategory(request modeltask.RequestUpdateCategory) (modeltask.ResponseStore, error) {
	err := validation.ValidateTaskUpdateCategory(request)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}
	entityTask := entity.Task{}
	copier.Copy(&entityTask, &request)
	update, err := s.repo.Update(entityTask)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}
	resp := modeltask.ResponseStore{}
	copier.Copy(&resp, &update)
	return resp, nil
}

func (s *Service) Update(request modeltask.RequestUpdate) (modeltask.ResponseStore, error) {
	// validator
	err := validation.ValidateTaskUpdate(request)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}

	entityTask := new(entity.Task)
	copier.Copy(&entityTask, &request)

	update, err := s.repo.Update(*entityTask)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}
	resp := new(modeltask.ResponseStore)
	copier.Copy(&resp, &update)

	return *resp, nil
}

func (s *Service) Create(request modeltask.Request) (modeltask.ResponseStore, error) {

	// validation
	err := validation.ValidateTaskCreate(request, s.repo)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}

	entityTask := new(entity.Task)
	copier.Copy(&entityTask, &request)

	entityTask.Status = false

	create, err := s.repo.Create(*entityTask)
	if err != nil {
		return modeltask.ResponseStore{}, err
	}

	resp := new(modeltask.ResponseStore)
	copier.Copy(&resp, &create)
	return *resp, nil
}

func (s *Service) Gets() ([]modeltask.ResponseGet, error) {
	tasks, err := s.repo.Gets()
	if err != nil {
		return nil, err
	}

	var resp []modeltask.ResponseGet

	copier.Copy(&resp, &tasks)

	return resp, nil
}

func New(repo repositorytask.RepositoryTask) ServiceTask {
	return &Service{repo: repo}
}