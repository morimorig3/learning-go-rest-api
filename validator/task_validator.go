package validator

import (
	"learning-go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	TaskValidate(task model.Task) error
}

type taskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

func (tv *taskValidator) TaskValidate(task model.Task) error {
	return validation.ValidateStruct(&task, validation.Field(
		&task.Title,
		validation.Required.Error("タイトルを入力してください"),
		validation.RuneLength(1, 10).Error("タイトルは10文字以内で入力してください"),
	))
}
