package main

import (
	"context"
	"geektime/app/goods/internal/conf"
	"geektime/pkg/appmanage"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	dbConf, httpConf, grpcConf, customerConf := conf.GenConf()
	app := initApp(dbConf, httpConf, grpcConf, customerConf)
	app.Register(&appmanage.RegisterInfo{
		Appid:   "goods:v1",
		AppName: "goods manager service",
	})
	app.Run(ctx, os.Interrupt)
	defer cancel()
}
