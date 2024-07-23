package rbac

import "context"

const (
	AppName = "rbac"
)

type Service interface {
	GetServiceAccountList(ctx context.Context, namespace, keyword string) ([]*ServiceAccountResponse, error)
	DeleteServiceAccount(ctx context.Context, namespace, name string) error
	CreateServiceAccount(ctx context.Context, req *ServiceAccountRequest) (*ServiceAccountResponse, error)
	GetRoleList(ctx context.Context, namespace, keyword string) ([]*RoleResponse, error)
	GetRoleDetail(ctx context.Context, namespace, name string) (*RoleRequest, error)
	CreateRole(ctx context.Context, req *RoleRequest) (*RoleResponse, error)
	DeleteRole(ctx context.Context, namespace, name string) error
	GetRoleBindingList(ctx context.Context, namespace, keyword string) ([]*RoleBindingResponse, error)
	GetRoleBindingDetail(ctx context.Context, namespace, name string) (*RoleBindingRequest, error)
	CreateRoleBinding(ctx context.Context, req *RoleBindingRequest) (*RoleBindingRequest, error)
	DeleteRoleBinding(ctx context.Context, namespace, name string) error
}
