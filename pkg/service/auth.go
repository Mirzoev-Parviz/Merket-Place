package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"market_place/models"
	"market_place/pkg/repository"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo repository.Authorization
}

type userTokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type merchTokenClaims struct {
	jwt.StandardClaims
	MerchantID int `json:"merchant_id"`
}

const (
	salt       = "ajsdijaskdasl122312klsdjka"
	signingKey = "kajsdljaskdja332$#"
)

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	user.IsActive = true
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (a *AuthService) GenerateUserToken(login, password string) (string, error) {
	user, err := a.repo.CheckUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &userTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}
func (a *AuthService) GenerateMerchantToken(login, password string) (string, error) {
	merch, err := a.repo.CheckMerch(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &merchTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		merch.ID,
	})

	return token.SignedString([]byte(signingKey))
}
func (a *AuthService) ParseUserToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &userTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")

		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*userTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *userTokenClaims")
	}

	return claims.UserID, nil
}

func (a *AuthService) ParseMerchantToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &merchTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")

		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*merchTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *merchTokenClaims")
	}

	return claims.MerchantID, nil
}
