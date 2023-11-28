package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	v2 "token_learn/api/token/v2"
	"token_learn/internal/biz"
	"token_learn/internal/conf"
)

const (
	// token:v2:${device_type}:${token_type}:${uid}
	rdKeyTokenV2 = "token:v2:%s:%s:%s"
)

type TokenV2Service struct {
	v2.UnimplementedTokenV2Server

	log         *log.Helper
	confTokenV2 *conf.TokenV2
	confData    *conf.Data
	redis       *redis.Client

	uc *biz.TokenUsecase
}

func NewTokenV2Service(logger log.Logger, confTokenV2 *conf.TokenV2, confData *conf.Data, uc *biz.TokenUsecase) *TokenV2Service {
	rdb := redis.NewClient(&redis.Options{
		Addr:     confData.Redis.Addr,
		Password: confData.Redis.Password,
		DB:       int(confTokenV2.Auth.Db),
	})
	return &TokenV2Service{
		log:         log.NewHelper(logger),
		confTokenV2: confTokenV2,
		confData:    confData,
		redis:       rdb,
		uc:          uc,
	}
}

// GenerateTokenV2 generates a new Token Token
func (s *TokenV2Service) GenerateTokenV2(ctx context.Context, req *v2.GenerateTokenV2Request) (*v2.GenerateTokenV2Response, error) {
	ctx = context.TODO()

	// 参数校验
	if _, ok := s.confTokenV2.Device.DeviceConfMap[req.DeviceType]; !ok {
		return nil, fmt.Errorf("unsupported device type")
	}

	refreshToken := generateToken(req.Uid, []byte(s.confTokenV2.Auth.Key))
	accessToken := generateToken(req.Uid, []byte(s.confTokenV2.Auth.Key))

	keyRefresh := fmt.Sprintf(rdKeyTokenV2, req.DeviceType, TokenTypeRefresh, req.Uid)
	results, err := s.redis.HGetAll(ctx, keyRefresh).Result()
	if err != nil {
		return nil, err
	}

	if int64(len(results)) >= s.confTokenV2.Device.DeviceConfMap[req.DeviceType] {
		return nil, fmt.Errorf("device limited max")
	}

	pipe := s.redis.TxPipeline()
	pipe.HSet(ctx, keyRefresh, req.DeviceId, refreshToken)
	pipe.Expire(ctx, keyRefresh, time.Second*time.Duration(s.confTokenV2.Auth.RefreshTokenExpiration))

	keyAccess := fmt.Sprintf(rdKeyTokenV2, req.DeviceType, TokenTypeAccess, req.Uid)
	pipe.HSet(ctx, keyAccess, req.DeviceId, accessToken)
	pipe.Expire(ctx, keyAccess, time.Second*time.Duration(s.confTokenV2.Auth.AccessTokenExpiration))

	if _, err := pipe.Exec(ctx); err != nil {
		// 回滚操作，删除已设置的 refresh token
		s.redis.Del(ctx, keyRefresh)
		s.redis.Del(ctx, keyAccess)
		s.log.Warnf("failed to set created token:%s", err.Error())
		return nil, err
	}
	return &v2.GenerateTokenV2Response{
		RefreshToken:           refreshToken,
		RefreshTokenExpiration: time.Now().Unix() + s.confTokenV2.Auth.RefreshTokenExpiration,
		AccessToken:            accessToken,
		AccessTokenExpiration:  time.Now().Unix() + s.confTokenV2.Auth.AccessTokenExpiration,
	}, nil
}

// DeleteTokenV2 generates a new Token Token
func (s *TokenV2Service) DeleteTokenV2(ctx context.Context, req *v2.DeleteTokenV2Request) (*v2.DeleteTokenV2Response, error) {
	ctx = context.TODO()

	if req.DeleteAll {
		pipe := s.redis.TxPipeline()
		for _, tokenType := range tokenList {
			for deviceType, _ := range s.confTokenV2.Device.DeviceConfMap {
				key := fmt.Sprintf(rdKeyTokenV2, deviceType, tokenType, req.Uid)
				pipe.Del(ctx, key)
			}
		}

		if _, err := pipe.Exec(ctx); err != nil {
			s.log.Warnf("failed to delete user %s all tokens:%s", req.Uid, err.Error())
			return nil, err
		}

		return &v2.DeleteTokenV2Response{
			Deleted: true,
		}, nil
	} else {
		key := fmt.Sprintf(rdKeyTokenV1, req.Type, req.Uid)
		result, err := s.redis.Del(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		if result < 0 {
			return &v2.DeleteTokenV2Response{
				Deleted: false,
			}, nil
		}
	}

	return &v2.DeleteTokenV2Response{
		Deleted: true,
	}, nil
}

// ValidateTokenV2 validates a Token Token
func (s *TokenV2Service) ValidateTokenV2(ctx context.Context, req *v2.ValidateTokenV2Request) (*v2.ValidateTokenV2Response, error) {
	ctx = context.TODO()

	key := fmt.Sprintf(rdKeyTokenV1, req.Type, req.Uid)
	result, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if result == "" || result != req.Token {
		return &v2.ValidateTokenV2Response{
			Valid: false,
		}, nil
	}

	return &v2.ValidateTokenV2Response{
		Valid: true,
	}, nil
}

// CreateAccessTokenV2 create an access token by refresh token
func (s *TokenV2Service) CreateAccessTokenV2(ctx context.Context, req *v2.CreateAccessTokenV2Request) (*v2.CreateAccessTokenV2Response, error) {
	ctx = context.TODO()

	key := fmt.Sprintf(rdKeyTokenV1, TokenTypeAccess, req.Uid)
	result, err := s.redis.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if result < 0 {
		return nil, fmt.Errorf("refresh token not found")
	}

	accessToken := generateToken(req.Uid, []byte(s.confTokenV2.Auth.Key))
	err = s.redis.Set(ctx, fmt.Sprintf(rdKeyTokenV1, TokenTypeAccess, req.Uid), accessToken, time.Second*time.Duration(s.confTokenV2.Auth.AccessTokenExpiration)).Err()
	if err != nil {
		return nil, err
	}
	return &v2.CreateAccessTokenV2Response{
		AccessToken:           accessToken,
		AccessTokenExpiration: time.Now().Unix() + s.confTokenV2.Auth.AccessTokenExpiration,
	}, nil
}
