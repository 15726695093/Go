//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"clock-in/app/worktime/job/internal/biz"
	"clock-in/app/worktime/job/internal/conf"
	"clock-in/app/worktime/job/internal/data"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*biz.WorktimeUsecase, error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet))
}
