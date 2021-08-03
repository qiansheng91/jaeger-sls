package sls_store

import (
	"context"
	slsSdk "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/jaegertracing/jaeger/model"
)

type slsSpanWriter struct {
	client   *slsSdk.Client
	instance slsTraceInstance
}

func (s slsSpanWriter) WriteSpan(ctx context.Context, span *model.Span) error {
	return nil
}
