//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"clock-in/app/worktime/service/internal/biz"
	"clock-in/app/worktime/service/internal/conf"
	"clock-in/app/worktime/service/internal/data"
	"clock-in/app/worktime/service/internal/server"
	"clock-in/app/worktime/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
