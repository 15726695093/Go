package service

import (
	"context"
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type CtxKey string

func ValidateToken() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token := tr.RequestHeader().Get("Authorization")
				user, _ := getIdFromToken(token)
				ctx = context.WithValue(ctx, "user", user)
			}
			return handler(ctx, req)
		}
	}
}

func getIdFromToken(token string) (int, error) {
	reply, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Split(string(reply), "|")[0])
}
