package server

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arkannsk/gql-example-service/config"
	"github.com/arkannsk/gql-example-service/server/graphql"
	auth_service "github.com/arkannsk/gql-example-service/services/auth-service"
	"github.com/go-pg/pg/v10"
	"net/http"
)

type server struct {
	listener *http.Server
}

func (s *server) ListenAndServe(addr string) error {
	s.listener.Addr = ":" + addr
	return s.listener.ListenAndServe()
}

func (s *server) Shutdown() {
	err := s.listener.Shutdown(context.Background())
	if err != nil {
		return
	}
}

func NewServer(db *pg.DB, authService *auth_service.AuthService) *server {
	srv := &server{
		listener: &http.Server{},
	}
	if config.Param.ENV == "dev" {
		http.Handle("/", playground.Handler("GraphQL", "/api/graphql/query"))
	}
	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("pong\n"))
	}))
	http.Handle("/api/graphql/query", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		realIP := r.Header.Get("X-Real-IP")
		ctx := r.Context()
		ctx = context.WithValue(ctx, "JWT", extractBearer(token))
		ctx = context.WithValue(ctx, "IP", realIP)
		handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: graphql.New(db, authService)})).
			ServeHTTP(w, r.WithContext(ctx))
	}))

	return srv
}
