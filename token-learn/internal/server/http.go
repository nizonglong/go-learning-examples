package server

import (
	v1Jwt "token_learn/api/jwt/v1"
	v1Token "token_learn/api/token/v1"
	v2Token "token_learn/api/token/v2"
	"token_learn/internal/conf"
	"token_learn/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(logger log.Logger, c *conf.Server, jwtService *service.JwtService, tokenService *service.TokenService, tokenV2Service *service.TokenV2Service) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1Jwt.RegisterJwtHTTPServer(srv, jwtService)
	v1Token.RegisterTokenHTTPServer(srv, tokenService)
	v2Token.RegisterTokenV2HTTPServer(srv, tokenV2Service)
	return srv
}
