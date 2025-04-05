package validator

import (
	"learning-go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (tv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user, validation.Field(
		&user.Email,
		validation.Required.Error("メールアドレスを入力してください"),
		validation.RuneLength(1, 50).Error("メールアドレスは50文字以内で入力してください"),
		is.Email.Error("メールアドレスを確認してください"),
	), validation.Field(
		&user.Password,
		validation.Required.Error("パスワードを入力してください"),
		validation.RuneLength(1, 16).Error("パスワードは16文字以内で入力してください"),
	))
}
