package log

import (
	"context"

	"github.com/chenquan/go-zero-demo/biztrace/trace"
	"github.com/zeromicro/go-zero/core/logx"
)

const traceIdKey = "bizTraceId"

func WithContext(ctx context.Context) logx.Logger {
	traceId, ok := trace.FromTraceId(ctx)
	if !ok {
		return logx.WithContext(ctx)
	}

	return logx.WithContext(ctx).WithFields(logx.Field(traceIdKey, traceId))
}
