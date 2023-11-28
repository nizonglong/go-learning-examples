package data

import (
	"context"

	"token_learn/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type jwtRepo struct {
	data *Data
	log  *log.Helper
}

// NewJwtRepo .
func NewJwtRepo(data *Data, logger log.Logger) biz.JwtRepo {
	return &jwtRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *jwtRepo) Save(ctx context.Context, g *biz.Jwt) (*biz.Jwt, error) {
	return g, nil
}

func (r *jwtRepo) Update(ctx context.Context, g *biz.Jwt) (*biz.Jwt, error) {
	return g, nil
}

func (r *jwtRepo) FindByID(context.Context, int64) (*biz.Jwt, error) {
	return nil, nil
}

func (r *jwtRepo) ListByHello(context.Context, string) ([]*biz.Jwt, error) {
	return nil, nil
}

func (r *jwtRepo) ListAll(context.Context) ([]*biz.Jwt, error) {
	return nil, nil
}
