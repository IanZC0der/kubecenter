package secret

import (
	"context"
	corev1 "k8s.io/api/core/v1"
)

const (
	AppName = "secret"
)

type Service interface {
	GetSecretsList(ctx context.Context, namespace string, keyword string) ([]*SecretResponse, error)
	GetSecretDetail(ctx context.Context, namespace string, name string) (*SecretResponse, error)
	DeleteSecret(ctx context.Context, namespace string, name string) error
	UpdateSecret(ctx context.Context, req *Secret) (*corev1.Secret, string, error)
}
