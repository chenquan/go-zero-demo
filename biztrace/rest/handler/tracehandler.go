package handler

import (
	"net/http"

	"github.com/chenquan/go-zero-demo/biztrace/trace"
)

func BizTraceHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			traceId := request.Header.Get("trace-id")
			if traceId == "" {
				next.ServeHTTP(writer, request)
				return
			}

			ctx := request.Context()
			ctx = trace.NewContext(ctx, traceId)
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}
