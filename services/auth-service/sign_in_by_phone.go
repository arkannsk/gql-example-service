package auth_service

import (
	"github.com/arkannsk/gql-example-service/db/models"
)

func (s *AuthService) SignInByPhoneCode(phone string, code string) (string, error) {
	userID, err := models.SignInByPhoneAuthCode(phone, code, s.phoneAuthMaxAttemptsCount, s.db)
	if err != nil {
		return "", err
	}
	token, err := s.CreateToken(userID)
	if err != nil {
		return "", err
	}
	return token, nil
}
