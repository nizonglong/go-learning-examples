package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Token struct {
	Hello string
}

type TokenRepo interface {
	Save(context.Context, *Token) (*Token, error)
	Update(context.Context, *Token) (*Token, error)
	FindByID(context.Context, int64) (*Token, error)
	ListByHello(context.Context, string) ([]*Token, error)
	ListAll(context.Context) ([]*Token, error)
}

type TokenUsecase struct {
	repo TokenRepo
	log  *log.Helper
}

func NewTokenUsecase(repo TokenRepo, logger log.Logger) *TokenUsecase {
	return &TokenUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TokenUsecase) SaveUserToken(ctx context.Context, userId, token string) (string, error) {
	uc.log.WithContext(ctx).Infof("GenerateToken: %v", userId)
	return token, nil
}
