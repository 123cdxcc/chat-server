//go:build wireinject
// +build wireinject

// wire.go
package cmd

import (
	"im-chat/internal/controller"
	"im-chat/internal/logic"

	"github.com/google/wire"
)

func InjectorApp() (*App, error) {
	wire.Build(
		NewApp,
		NewNoAuthServer,
		NewAuthServer,
		controller.AuthProviderSet,
		controller.NoAuthProviderSet,
		logic.ProviderSet,
	)
	return nil, nil
}
