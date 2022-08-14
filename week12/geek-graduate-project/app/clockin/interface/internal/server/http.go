package server

import (
	v1 "clock-in/api/clockin/interface/v1"
	"clock-in/app/clockin/interface/internal/conf"
	"clock-in/app/clockin/interface/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Server,
	clockin *service.ClockInService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			selector.Server(
				service.ValidateToken(),
			).
				Path(
					"/api.clockin.interface.v1.ClockinInterface/ClockinOnWork",
					"/api.clockin.interface.v1.ClockinInterface/ClockinOffWork",
				).
				Build(),
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
	v1.RegisterClockinInterfaceHTTPServer(srv, clockin)
	return srv
}
