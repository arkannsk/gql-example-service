// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"github.com/arkannsk/gql-example-service/db/models"
)

type SignInOrErrorPayload interface {
	IsSignInOrErrorPayload()
}

type ErrorPayload struct {
	Message string `json:"message"`
}

func (ErrorPayload) IsSignInOrErrorPayload() {}

type RequestSignInCodeInput struct {
	Phone string `json:"phone"`
}

type SignInByCodeInput struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

type SignInPayload struct {
	Token  string  `json:"token"`
	Viewer *Viewer `json:"viewer"`
}

func (SignInPayload) IsSignInOrErrorPayload() {}

type Viewer struct {
	User *models.User `json:"user"`
}
