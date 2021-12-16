package auth_service

import (
	"github.com/arkannsk/gql-example-service/db/models"
)

func (s *AuthService) CreateRequestSignInByCode(phone string) error {
	return models.CreateRequestSignInByCode(phone, genNumCode(4), s.phoneAuthCodeTTL, s.db)
}
