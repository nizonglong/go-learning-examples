package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewJwtService, NewTokenService, NewTokenV2Service)
