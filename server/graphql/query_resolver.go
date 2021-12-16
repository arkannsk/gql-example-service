package graphql

import (
	"context"
	"github.com/arkannsk/gql-example-service/db/models"
	gqlmodel "github.com/arkannsk/gql-example-service/server/graphql/model"
	log "github.com/sirupsen/logrus"
)

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	res, err := models.GetRandomProducts(10, r.db)
	if err != nil {
		log.Error(err)
		return nil, nil
	}
	return res, nil
}

func (r *queryResolver) Viewer(ctx context.Context) (*gqlmodel.Viewer, error) {
	token := ctx.Value("JWT").(string)
	if len(token) == 0 {
		return &gqlmodel.Viewer{}, nil
	}
	user, err := r.authService.GetUserByJWTToken(token)
	if err != nil {
		log.Error(err)
		return &gqlmodel.Viewer{}, nil
	}
	return &gqlmodel.Viewer{User: user}, nil
}
