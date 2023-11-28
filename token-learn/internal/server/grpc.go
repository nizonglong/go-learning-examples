package server

import (
	"github.com/go-kratos/kratos/v2/log"

	v1Jwt "token_learn/api/jwt/v1"
	v1Token "token_learn/api/token/v1"
	v2Token "token_learn/api/token/v2"
	"token_learn/internal/conf"
	"token_learn/internal/service"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(logger log.Logger, c *conf.Server, jwtService *service.JwtService, tokenService *service.TokenService, tokenV2Service *service.TokenV2Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1Jwt.RegisterJwtServer(srv, jwtService)
	v1Token.RegisterTokenServer(srv, tokenService)
	v2Token.RegisterTokenV2Server(srv, tokenV2Service)
	return srv
}
