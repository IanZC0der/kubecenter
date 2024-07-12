package configmap

import "context"

const (
	AppName = "configmap"
)

type Service interface {
	GetConfigMapList(ctx context.Context, namespace string, keyword string) ([]*ConfigMapResponse, error)
	GetConfigMapDetail(ctx context.Context, namespace string, name string) (*ConfigMapResponse, error)
}
