package auth_service

import (
	"crypto/rsa"
	"github.com/go-pg/pg/v10"
	"log"
)

type AuthService struct {
	db                        *pg.DB
	phoneAuthCodeTTL          int // in seconds
	phoneAuthMaxAttemptsCount int // in seconds
	jwtPublicKey              *rsa.PublicKey
	jwtPrivateKey             *rsa.PrivateKey
	JwtTokenTTL               int // in minutes
}

func NewAuthService(
	db *pg.DB,
	phoneAuthCodeTTL int,
	phoneAuthMaxAttemptsCount int,
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
		phoneAuthMaxAttemptsCount: phoneAuthMaxAttemptsCount,
		jwtPublicKey:              publicKey,
		jwtPrivateKey:             privateKey,
		JwtTokenTTL:               jwtTokenTTL,
	}
}
