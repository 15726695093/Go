package server

import (
	"context"
	v1 "geektime/api/goods/v1"
	"geektime/app/goods/internal/conf"
	"geektime/app/goods/internal/service"
	"geektime/pkg/appmanage"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func (g *GrpcServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		g.server.Stop()
	}()
	return g.server.Serve(g.listener)
}

func NewGrpcServer(service *service.GoodsService, config *conf.GrpcConf) appmanage.GrpcServer {
	server := new(GrpcServer)
	lis, err := net.Listen("tcp", config.Addr())
	server.listener = lis
	if err != nil {
		panic(err.Error())
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	v1.RegisterGoodsServiceServer(grpcServer, service)
	server.server = grpcServer
	return server
}
