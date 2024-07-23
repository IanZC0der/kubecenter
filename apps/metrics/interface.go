package metrics

import "context"

const (
	AppName = "metrics"
)

type Service interface {
	GetClusterInfo(ctx context.Context) []*MetricsItem
}
