package graphql

import (
	"context"
	gqlmodel "github.com/arkannsk/gql-example-service/server/graphql/model"
	"log"
)

func (r *mutationResolver) RequestSignInCode(ctx context.Context, input gqlmodel.RequestSignInCodeInput) (*gqlmodel.ErrorPayload, error) {
	err := r.authService.CreateRequestSignInByCode(input.Phone)
	if err != nil {
		log.Print(err)
		return &gqlmodel.ErrorPayload{Message: err.Error()}, nil
	}
	return nil, nil
}

func (r *mutationResolver) SignInByCode(ctx context.Context, input gqlmodel.SignInByCodeInput) (gqlmodel.SignInOrErrorPayload, error) {
	token, err := r.authService.SignInByPhoneCode(input.Phone, input.Code)
	if err != nil {
		return gqlmodel.ErrorPayload{
			Message: err.Error(),
		}, nil
	}
	user, err := r.authService.GetUserByJWTToken(token)
	if err != nil {
		return gqlmodel.ErrorPayload{
			Message: err.Error(),
		}, nil
	}
	return gqlmodel.SignInPayload{
		Token:  token,
		Viewer: &gqlmodel.Viewer{User: user},
	}, nil
}
