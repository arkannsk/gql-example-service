package graphql

import (
	auth_service "github.com/arkannsk/gql-example-service/services/auth-service"
	"github.com/go-pg/pg/v10"
)

type Resolver struct {
	db          *pg.DB
	authService *auth_service.AuthService
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

type Option func(*Resolver)

func New(db *pg.DB, service *auth_service.AuthService) ResolverRoot {
	return &Resolver{
		db:          db,
		authService: service,
	}
}
