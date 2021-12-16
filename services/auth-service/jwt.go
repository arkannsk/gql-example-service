package auth_service

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/arkannsk/gql-example-service/db/models"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-pg/pg/v10"
	"os"
	"strconv"
	"time"
)

type Claims struct {
	UserID   string
	Audience string
}

type Token struct {
	str    string
	Claims Claims
	TTL    time.Duration
}

func (t Token) String() string {
	return t.str
}

func (s *AuthService) Parse(tokenString string) (*Token, error) {
	if tokenString == "" {
		return nil, nil
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtPublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil
	}

	result := &Token{
		str: tokenString,
		Claims: Claims{
			UserID: fmt.Sprintf("%s", claims["user_id"]),
		},
		TTL: time.Duration(int64(claims["exp"].(float64))-time.Now().Unix()) * time.Second,
	}

	return result, nil
}

func (s *AuthService) CreateToken(userID int) (string, error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["exp"] = now.Add(time.Minute * time.Duration(s.JwtTokenTTL)).Unix()
	claims["iat"] = now.Unix()
	claims["user_id"] = fmt.Sprintf("%d", userID)
	t.Claims = claims

	return t.SignedString(s.jwtPrivateKey)
}

func (s *AuthService) GetUserByJWTToken(tokenString string) (*models.User, error) {
	token, err := s.Parse(tokenString)
	if err != nil {
		return nil, err
	}
	userID, _ := strconv.ParseInt(token.Claims.UserID, 10, 64)
	user, err := models.GetUserByCriteria(models.UserCriteria{ID: int(userID)}, s.db)
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}
	return user, nil
}

func ReadPublicKey(path string) (*rsa.PublicKey, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(dat)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func ReadPrivateKey(path string) (*rsa.PrivateKey, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPrivateKeyFromPEM(dat)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
