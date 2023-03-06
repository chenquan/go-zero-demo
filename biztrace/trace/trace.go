package trace

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go.opentelemetry.io/otel/baggage"
)

const traceIdKey = "biz-trace-id"

func NewContext(ctx context.Context, traceId string) context.Context {
	logger := logx.WithContext(ctx).WithFields(logx.Field(traceId, traceId))

	bg := baggage.FromContext(ctx)
	member, err := baggage.NewMember(traceIdKey, traceId)
	if err != nil {
		logger.Error(err)
		return ctx
	}

	bg, err = bg.SetMember(member)
	if err != nil {
		logger.Error(err)
		return ctx
	}

	ctx = baggage.ContextWithBaggage(ctx, bg)

	return ctx
}

func FromTraceId(ctx context.Context) (string, bool) {
	bg := baggage.FromContext(ctx)
	member := bg.Member(traceIdKey)
	return member.Value(), member.Key() != ""
}
