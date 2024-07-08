package pods

import (
	"context"
)

const (
	AppName = "pods"
)

type Service interface {
	GetPods(ctx context.Context) (*Pods, error)
}
