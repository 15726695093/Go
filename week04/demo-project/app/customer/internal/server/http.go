package server

import (
	"context"
	v1 "geektime/api/customer/v1"
	"geektime/app/customer/internal/conf"
	"geektime/app/customer/internal/service"
	"geektime/pkg/appmanage"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
}

func (s *HttpServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.server.Shutdown(ctx)
	}()
	return s.server.ListenAndServe()
}

func NewHttpServer(service *service.CustomerService, config *conf.HttpConf) appmanage.HttpServer {
	server := new(HttpServer)
	engine := gin.Default()
	v1.RegisterCustomerHttpServer(engine, service)
	server.server = &http.Server{
		Addr:    config.Addr(),
		Handler: engine,
	}
	return server
}
