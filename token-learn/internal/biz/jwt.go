package biz

import (
	"context"

	v1 "token_learn/api/jwt/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type Jwt struct {
	Hello string
}

type JwtRepo interface {
	Save(context.Context, *Jwt) (*Jwt, error)
	Update(context.Context, *Jwt) (*Jwt, error)
	FindByID(context.Context, int64) (*Jwt, error)
	ListByHello(context.Context, string) ([]*Jwt, error)
	ListAll(context.Context) ([]*Jwt, error)
}

type JwtUsecase struct {
	repo JwtRepo
	log  *log.Helper
}

func NewJwtUsecase(repo JwtRepo, logger log.Logger) *JwtUsecase {
	return &JwtUsecase{repo: repo, log: log.NewHelper(logger)}
}
