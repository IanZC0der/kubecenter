package configmap

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

const (
	AppName = "configmap"
)

type Service interface {
	GetConfigMapList(ctx context.Context, namespace string, keyword string) ([]*ConfigMapResponse, error)
	GetConfigMapDetail(ctx context.Context, namespace string, name string) (*ConfigMapResponse, error)
	CreateOrUpdateConfigMap(ctx context.Context, configMap *ConfigMap) (*corev1.ConfigMap, string, error)
	DeleteConfigMap(ctx context.Context, namespace string, name string) error
}
