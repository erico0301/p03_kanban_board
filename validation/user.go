package validation

import (
	"errors"

	"p03_kanban_board/model/modeluser"
	"p03_kanban_board/repository/repositoryuser"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func isEmailExist(repo repositoryuser.RepositoryUser) validation.RuleFunc {
	return func(value interface{}) error {
		email, ok := value.(string)
		if !ok {
			return errors.New("invalid email address")
		}

		return repo.IsEmailExist(email)
	}
}

func ValidateUserCreate(data modeluser.Request, repo repositoryuser.RepositoryUser) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email, validation.By(isEmailExist(repo))),
		"fullname": validation.Validate(data.Fullname, validation.Required),
		"password": validation.Validate(data.Password, validation.Required, validation.Length(8, 20)),
	}.Filter()
}

func ValidateUserLogin(data modeluser.RequestLogin) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email),
		"password": validation.Validate(data.Password, validation.Required, validation.Length(8, 20).Error("invalid email or password")),
	}.Filter()
}

func ValidateUserUpdate(data modeluser.Request) error {
	return validation.Errors{
		"email":    validation.Validate(data.Email, validation.Required, is.Email),
		"fullname": validation.Validate(data.Fullname, validation.Required),
	}.Filter()
}
