package main

import (
	"clock-in/app/worktime/job/internal/conf"
	"clock-in/app/worktime/job/internal/pkg/mq"
	"context"
	"flag"
	stdlog "log"
	"os"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	uc, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	consumer := mq.NewConsumer(bc.Data)
	go func() {
		for m := range consumer {
			err := uc.CaculateWorktime(context.Background(), string(m.Body))
			if err != nil {
				stdlog.Printf("caculate worktime failed: reason:%v", err.Error())
			}
		}
	}()
	<-forever
}
