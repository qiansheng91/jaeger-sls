package sls_store

import (
	"context"
	slsSdk "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/jaegertracing/jaeger/model"
	"time"
)

type slsDependencyReader struct {
	client   *slsSdk.Client
	instance slsTraceInstance
}

func (s slsDependencyReader) GetDependencies(ctx context.Context, endTs time.Time, lookback time.Duration) ([]model.DependencyLink, error) {
	return nil, nil
}
