package service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"

	v1 "token_learn/api/jwt/v1"
	"token_learn/internal/biz"
	"token_learn/internal/conf"
)

type JwtService struct {
	v1.UnimplementedJwtServer

	conf *conf.Jwt
	uc   *biz.JwtUsecase
}

func NewJwtService(conf *conf.Jwt, uc *biz.JwtUsecase) *JwtService {
	return &JwtService{
		conf: conf,
		uc:   uc,
	}
}

// GenerateJwt generates a new JWT jwt
func (s *JwtService) GenerateJwt(ctx context.Context, req *v1.GenerateJwtRequest) (*v1.GenerateJwtResponse, error) {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Duration(s.conf.ExpireTime)).Unix(), // Jwt expiration time
		"username": req.Username,
		"email":    req.Email,
		"phone":    req.Phone,
		// Add other claims as needed
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(s.conf.Key)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &v1.GenerateJwtResponse{
		Token: signedToken,
	}, nil
}

// ValidateJwt validates a JWT jwt
func (s *JwtService) ValidateJwt(ctx context.Context, req *v1.ValidateJwtRequest) (*v1.ValidateJwtResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.conf.Key), nil
	})

	if err != nil {
		return nil, err
	}

	return &v1.ValidateJwtResponse{
		Valid: token.Valid,
	}, nil
}
