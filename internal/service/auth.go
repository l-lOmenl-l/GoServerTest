package service

import (
	"crypto/sha1"
	"errors"
	"example/web-servise-gin/internal/domain"
	"example/web-servise-gin/internal/repository"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type tokenClaims struct {
	jwt.StandardClaims
	Place       string `json:"place"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	FullName    string `json:"fullName"`
}

type tokenClaimsSRS struct {
	jwt.StandardClaims
	domain.User
}

type AuthService struct {
	repo repository.Authorization
}

const (
	salt       = "greafgho3234wef"
	signingKey = "wefw34efwsgasdlj342kghlswei"
	tokenTTL   = 1 * time.Hour
)

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.SignUp) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(inf domain.SignIn) (string, error) {

	if inf.Place == "SRS" {
		user, err := s.repo.GetUser(inf.Login, generatePasswordHash(inf.Password))
		if err != nil {
			return "", err
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaimsSRS{jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
			user,
		},
		)
		strToken, err := token.SignedString([]byte(signingKey))
		if err != nil {
			return "", err
		}

		err = s.repo.SaveToken(strToken)
		if err != nil {
			return "", err
		}

		return strToken, nil

	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
			inf.Place,
			inf.Address,
			inf.PhoneNumber,
			inf.FIO,
		},
		)
		strToken, err := token.SignedString([]byte(signingKey))
		if err != nil {
			return "", err
		}

		err = s.repo.SaveToken(strToken)
		if err != nil {
			return "", err
		}

		return strToken, nil
	}

}

func (s *AuthService) ParseToken(accessToken string) (string, error) {

	err := s.repo.IdentifyToken(accessToken)
	if err != nil {
		return "", errors.New("token don't find")
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.FullName, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
