package metadata

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func Metadata() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if _, ok := transport.FromServerContext(ctx); ok {
				log.Println("get in")
				defer func() {
				}()
			}
			ctx = metadata.AppendToClientContext(ctx, "x-md-global-extra", "2233")
			return handler(ctx, req)
		}
	}
}
