package validator

import (
	"user-service/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user domain.User) error
	SignupRequestValidate(user domain.SignupRequest) error
	LoginRequestValidate(user domain.LoginRequest) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user domain.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Username,
			validation.Required.Error("username is required"),
			validation.RuneLength(6, 20).Error("limited min 6 max 20 char"),
		),
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limited min 8 max 30 char"),
		),
	)
}

func (uv *userValidator) SignupRequestValidate(req domain.SignupRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Username,
			validation.Required.Error("username is required"),
			validation.RuneLength(6, 20).Error("limited min 6 max 20 char"),
		),
		validation.Field(
			&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limited min 8 max 30 char"),
		),
	)
}

func (uv *userValidator) LoginRequestValidate(req domain.LoginRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&req.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 30).Error("limited min 8 max 30 char"),
		),
	)
}
