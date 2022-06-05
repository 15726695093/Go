//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"geektime/app/customer/internal/biz"
	"geektime/app/customer/internal/conf"
	"geektime/app/customer/internal/data"
	"geektime/app/customer/internal/server"
	"geektime/app/customer/internal/service"
	"geektime/pkg/appmanage"

	"github.com/google/wire"
)

func initApp(
	db *conf.ConfDB,
	http *conf.HttpConf,
	grpc *conf.GrpcConf,
) *appmanage.AppManage {
	panic(wire.Build(
		server.ProvideSet,
		data.ProvideSet,
		service.ProvideSet,
		biz.ProvideSet,
		appmanage.NewAppManage,
	))
}
