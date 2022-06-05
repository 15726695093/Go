//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"geektime/app/goods/internal/biz"
	"geektime/app/goods/internal/conf"
	"geektime/app/goods/internal/data"
	"geektime/app/goods/internal/server"
	"geektime/app/goods/internal/service"
	"geektime/pkg/appmanage"

	"github.com/google/wire"
)

func initApp(
	db *conf.ConfDB,
	http *conf.HttpConf,
	grpc *conf.GrpcConf,
	customer *conf.Customer,
) *appmanage.AppManage {
	panic(wire.Build(
		server.ProvideSet,
		data.ProvideSet,
		service.ProvideSet,
		biz.ProvideSet,
		appmanage.NewAppManage,
	))
}
