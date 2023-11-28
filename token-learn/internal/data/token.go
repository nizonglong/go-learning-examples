package data

import (
	"context"

	"token_learn/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type tokenRepo struct {
	data *Data
	log  *log.Helper
}

// NewTokenRepo .
func NewTokenRepo(data *Data, logger log.Logger) biz.TokenRepo {
	return &tokenRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *tokenRepo) Save(ctx context.Context, g *biz.Token) (*biz.Token, error) {
	return g, nil
}

func (r *tokenRepo) Update(ctx context.Context, g *biz.Token) (*biz.Token, error) {
	return g, nil
}

func (r *tokenRepo) FindByID(context.Context, int64) (*biz.Token, error) {
	return nil, nil
}

func (r *tokenRepo) ListByHello(context.Context, string) ([]*biz.Token, error) {
	return nil, nil
}

func (r *tokenRepo) ListAll(context.Context) ([]*biz.Token, error) {
	return nil, nil
}
