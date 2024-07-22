package rbac

import "context"

const (
	AppName = "rbac"
)

type Service interface {
	GetServiceAccountList(ctx context.Context, namespace, keyword string) ([]*ServiceAccountResponse, error)
	DeleteServiceAccount(ctx context.Context, namespace, name string) error
	CreateServiceAccount(ctx context.Context, req *ServiceAccountRequest) (*ServiceAccountResponse, error)
}
