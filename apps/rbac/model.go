package rbac

import (
	"github.com/IanZC0der/kubecenter/util"
	rbacv1 "k8s.io/api/rbac/v1"
)

type RequestBase struct {
	Name      string           `json:"name"`
	Namespace string           `json:"namespace"`
	Labels    []*util.ListItem `json:"labels"`
}

type ResponseBase struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Age       int64  `json:"age"`
}

type ServiceAccountRequest struct {
	*RequestBase
}

type ServiceAccountResponse struct {
	*ResponseBase
}

type RoleRequest struct {
	*RequestBase
	Rules []rbacv1.PolicyRule `json:"rules"`
}

type RoleResponse struct {
	*ResponseBase
}

type CreateRoleBindingRequest struct {
	*RequestBase
	Subjects []*ServiceAccountRequest `json:"subjects"`
	RoleRef  string                   `json:"roleRef"`
}

type RoleBindingResponse struct {
	*ResponseBase
}

func NewCreateServiceAccountRequest() *ServiceAccountRequest {
	return &ServiceAccountRequest{
		RequestBase: &RequestBase{
			Labels: make([]*util.ListItem, 0),
		},
	}
}

func NewRoleRequest() *RoleRequest {
	return &RoleRequest{
		RequestBase: &RequestBase{
			Labels: make([]*util.ListItem, 0),
		},
		Rules: make([]rbacv1.PolicyRule, 0),
	}
}
