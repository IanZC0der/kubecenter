package metrics

import "context"

const (
	AppName = "metrics"
)

type Service interface {
	GetClusterInfo(ctx context.Context) []*MetricsItem
	GetResourceInfo(ctx context.Context) []*MetricsItem
	GetClusterUsageInfo(ctx context.Context) []*MetricsItem
}
