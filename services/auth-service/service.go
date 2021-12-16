package auth_service

import (
	"crypto/rsa"
	"github.com/go-pg/pg/v10"
	"log"
)

type AuthService struct {
	db                        *pg.DB
	phoneAuthCodeTTL          int // in seconds
	phoneAuthMaxWrongAttempts int // in seconds
	jwtPublicKey              *rsa.PublicKey
	jwtPrivateKey             *rsa.PrivateKey
	JwtTokenTTL               int // in minutes
}

func NewAuthService(
	db *pg.DB,
	phoneAuthCodeTTL int,
	phoneAuthMaxWrongAttempts int,
	jwtTokenTTL int,
	jwtPublicKeyPath string,
	jwtPrivateKeyPath string) *AuthService {
	publicKey, err := ReadPublicKey(jwtPublicKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := ReadPrivateKey(jwtPrivateKeyPath)
	if err != nil {
		log.Fatal(err)
	}
	return &AuthService{
		db:                        db,
		phoneAuthCodeTTL:          phoneAuthCodeTTL,
		phoneAuthMaxWrongAttempts: phoneAuthMaxWrongAttempts,
		jwtPublicKey:              publicKey,
		jwtPrivateKey:             privateKey,
		JwtTokenTTL:               jwtTokenTTL,
	}
}

func (s *AuthService) GeneratePhoneAuthCode() string {
	return ""
}
