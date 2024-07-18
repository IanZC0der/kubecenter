package k8sservice

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

const (
	AppName = "K8sService"
)

type Service interface {
	GetSvcList(ctx context.Context, namespace string, keyword string) ([]*corev1.Service, error)
	GetSvcDetail(ctx context.Context, namespace string, name string) (*corev1.Service, error)
	CreateOrUpdateSvc(ctx context.Context, req *K8SService) (*corev1.Service, string, error)
	DeleteSvc(ctx context.Context, namespace, name string) error
}
