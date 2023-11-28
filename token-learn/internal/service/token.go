package service

import (
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	v1 "token_learn/api/token/v1"
	"token_learn/internal/biz"
	"token_learn/internal/conf"
)

const (
	// token:${token_type}:${uid}
	rdKeyTokenV1     = "token:v1:%s:%s"
	TokenTypeRefresh = "refresh"
	TokenTypeAccess  = "access"
)

var (
	tokenList = []string{
		TokenTypeRefresh,
		TokenTypeAccess,
	}
)

type TokenService struct {
	v1.UnimplementedTokenServer

	log       *log.Helper
	confToken *conf.Token
	confData  *conf.Data
	redis     *redis.Client

	uc *biz.TokenUsecase
}

func NewTokenService(logger log.Logger, confToken *conf.Token, confData *conf.Data, uc *biz.TokenUsecase) *TokenService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     confData.Redis.Addr,
		Password: confData.Redis.Password,
		DB:       int(confToken.Auth.Db),
	})

	return &TokenService{
		log:       log.NewHelper(logger),
		confToken: confToken,
		confData:  confData,
		redis:     rdb,
		uc:        uc,
	}
}

func generateToken(userID string, secretKey []byte) string {
	// 生成消息，可以是用户ID或其他需要包含的信息
	message := []byte(userID)

	// 生成随机盐
	salt := generateSalt()

	// 将盐和消息结合
	data := append(message, salt...)

	// 使用 HMAC-SHA256 生成签名
	hash := hmac.New(sha256.New, secretKey)
	hash.Write(data)
	signature := hash.Sum(nil)

	// 将签名、盐和消息结合编码为 base64 字符串作为最终令牌
	token := base64.URLEncoding.EncodeToString(signature) + "." + base64.URLEncoding.EncodeToString(salt) + "." +
		base64.URLEncoding.EncodeToString(message)
	return token
}

func generateSalt() []byte {
	// 使用安全的随机源生成盐
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

// GenerateToken generates a new Token Token
func (s *TokenService) GenerateToken(ctx context.Context, req *v1.GenerateTokenRequest) (*v1.GenerateTokenResponse, error) {
	ctx = context.TODO()
	refreshToken := generateToken(req.Uid, []byte(s.confToken.Auth.Key))
	accessToken := generateToken(req.Uid, []byte(s.confToken.Auth.Key))

	pipe := s.redis.TxPipeline()
	keyRefresh := fmt.Sprintf(rdKeyTokenV1, TokenTypeRefresh, req.Uid)
	keyAccess := fmt.Sprintf(rdKeyTokenV1, TokenTypeAccess, req.Uid)
	pipe.Set(ctx, keyRefresh, refreshToken, time.Second*time.Duration(s.confToken.Auth.RefreshTokenExpiration))
	pipe.Set(ctx, keyAccess, accessToken, time.Second*time.Duration(s.confToken.Auth.AccessTokenExpiration))
	if _, err := pipe.Exec(ctx); err != nil {
		// 回滚操作，删除已设置的 key
		s.redis.Del(ctx, keyRefresh)
		s.redis.Del(ctx, keyAccess)
		s.log.Warnf("failed to set created token:%s", err.Error())
		return nil, err
	}

	return &v1.GenerateTokenResponse{
		RefreshToken:           refreshToken,
		RefreshTokenExpiration: time.Now().Unix() + s.confToken.Auth.RefreshTokenExpiration,
		AccessToken:            accessToken,
		AccessTokenExpiration:  time.Now().Unix() + s.confToken.Auth.AccessTokenExpiration,
	}, nil
}

// DeleteToken generates a new Token Token
func (s *TokenService) DeleteToken(ctx context.Context, req *v1.DeleteTokenRequest) (*v1.DeleteTokenResponse, error) {
	ctx = context.TODO()

	if req.DeleteAll {
		for _, ty := range tokenList {
			key := fmt.Sprintf(rdKeyTokenV1, ty, req.Uid)
			_, err := s.redis.Del(ctx, key).Result()
			if err != nil {
				return nil, err
			}
		}
		return &v1.DeleteTokenResponse{
			Deleted: true,
		}, nil
	} else {
		key := fmt.Sprintf(rdKeyTokenV1, req.Type, req.Uid)
		result, err := s.redis.Del(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		if result < 0 {
			return &v1.DeleteTokenResponse{
				Deleted: false,
			}, nil
		}
	}

	return &v1.DeleteTokenResponse{
		Deleted: true,
	}, nil
}

// ValidateToken validates a Token Token
func (s *TokenService) ValidateToken(ctx context.Context, req *v1.ValidateTokenRequest) (*v1.ValidateTokenResponse, error) {
	ctx = context.TODO()

	key := fmt.Sprintf(rdKeyTokenV1, req.Type, req.Uid)
	result, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if result == "" || result != req.Token {
		return &v1.ValidateTokenResponse{
			Valid: false,
		}, nil
	}

	return &v1.ValidateTokenResponse{
		Valid: true,
	}, nil
}

// CreateAccessToken create an access token by refresh token
func (s *TokenService) CreateAccessToken(ctx context.Context, req *v1.CreateAccessTokenRequest) (*v1.CreateAccessTokenResponse, error) {
	ctx = context.TODO()

	key := fmt.Sprintf(rdKeyTokenV1, TokenTypeAccess, req.Uid)
	result, err := s.redis.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	if result < 0 {
		return nil, fmt.Errorf("refresh token not found")
	}

	accessToken := generateToken(req.Uid, []byte(s.confToken.Auth.Key))
	err = s.redis.Set(ctx, fmt.Sprintf(rdKeyTokenV1, TokenTypeAccess, req.Uid), accessToken, time.Second*time.Duration(s.confToken.Auth.AccessTokenExpiration)).Err()
	if err != nil {
		return nil, err
	}
	return &v1.CreateAccessTokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiration: time.Now().Unix() + s.confToken.Auth.AccessTokenExpiration,
	}, nil
}
