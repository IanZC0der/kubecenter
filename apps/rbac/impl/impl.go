package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/rbac"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func init() {
	ioc.DefaultControllerContainer().Register(&RBACServiceImpl{})
}

var _ rbac.Service = &RBACServiceImpl{}

type RBACServiceImpl struct{}

func (s *RBACServiceImpl) Init() error {
	return nil
}

func (s *RBACServiceImpl) Name() string {
	return rbac.AppName
}

func (s *RBACServiceImpl) GetServiceAccountList(ctx context.Context, namespace, keyword string) ([]*rbac.ServiceAccountResponse, error) {
	l, err := global.KubeConfigSet.CoreV1().ServiceAccounts(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	res := make([]*rbac.ServiceAccountResponse, 0)
	for _, item := range l.Items {
		res = append(res, &rbac.ServiceAccountResponse{
			ResponseBase: &rbac.ResponseBase{
				Name:      item.Name,
				Namespace: item.Namespace,
				Age:       item.CreationTimestamp.Unix(),
			},
		})
	}
	return res, nil
}
func (s *RBACServiceImpl) DeleteServiceAccount(ctx context.Context, namespace, name string) error {
	err := global.KubeConfigSet.CoreV1().ServiceAccounts(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
func (s *RBACServiceImpl) CreateServiceAccount(ctx context.Context, req *rbac.ServiceAccountRequest) (*rbac.ServiceAccountResponse, error) {
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels:    util.MapConverter(req.Labels),
		},
	}

	sa, err := global.KubeConfigSet.CoreV1().ServiceAccounts(req.Namespace).Create(ctx, serviceAccount, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	res := &rbac.ServiceAccountResponse{
		ResponseBase: &rbac.ResponseBase{
			Name:      sa.Name,
			Namespace: sa.Namespace,
			Age:       sa.CreationTimestamp.Unix(),
		},
	}
	return res, nil
}
func (s *RBACServiceImpl) GetRoleList(ctx context.Context, namespace, keyword string) ([]*rbac.RoleResponse, error) {
	res := make([]*rbac.RoleResponse, 0)
	if namespace != "" {
		list, err := global.KubeConfigSet.RbacV1().Roles(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, item := range list.Items {
			if !strings.Contains(item.Name, keyword) {
				continue
			}
			res = append(res, &rbac.RoleResponse{
				ResponseBase: &rbac.ResponseBase{
					Name:      item.Name,
					Namespace: item.Namespace,
					Age:       item.CreationTimestamp.Unix(),
				},
			})
		}
	} else {
		clusterRoleList, err := global.KubeConfigSet.RbacV1().ClusterRoles().List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, item := range clusterRoleList.Items {
			if !strings.Contains(item.Name, keyword) {
				continue
			}
			res = append(res, &rbac.RoleResponse{
				ResponseBase: &rbac.ResponseBase{
					Name:      item.Name,
					Namespace: item.Namespace,
					Age:       item.CreationTimestamp.Unix(),
				},
			})
		}

	}
	return res, nil
}
func (s *RBACServiceImpl) GetRoleDetail(ctx context.Context, namespace, name string) (*rbac.RoleRequest, error) {
	res := rbac.NewRoleRequest()

	if namespace != "" {
		role, err := global.KubeConfigSet.RbacV1().Roles(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		res.Name = role.Name
		res.Namespace = role.Namespace
		res.Labels = util.ListConverter(role.Labels)
		res.Rules = role.Rules

	} else {
		role, err := global.KubeConfigSet.RbacV1().ClusterRoles().Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = role.Name
		res.Namespace = role.Namespace
		res.Labels = util.ListConverter(role.Labels)
		res.Rules = role.Rules
	}
	return res, nil
}
func (s *RBACServiceImpl) CreateRole(ctx context.Context, req *rbac.RoleRequest) (*rbac.RoleResponse, error) {
	// if namespace is not empty, create role in namespace, else in the cluster
	res := &rbac.RoleResponse{
		ResponseBase: &rbac.ResponseBase{},
	}
	if req.Namespace != "" {
		nsRole := &rbacv1.Role{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name,
				Namespace: req.Namespace,
				Labels:    util.MapConverter(req.Labels),
			},
			Rules: req.Rules,
		}
		created, err := global.KubeConfigSet.RbacV1().Roles(nsRole.Namespace).Create(ctx, nsRole, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = created.Name
		res.Namespace = created.Namespace
		res.Age = created.CreationTimestamp.Unix()
	} else {
		clusterRole := &rbacv1.ClusterRole{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name,
				Namespace: req.Namespace,
				Labels:    util.MapConverter(req.Labels),
			},
			Rules: req.Rules,
		}
		created, err := global.KubeConfigSet.RbacV1().ClusterRoles().Create(ctx, clusterRole, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = created.Name
		res.Namespace = created.Namespace
		res.Age = created.CreationTimestamp.Unix()
	}
	return res, nil
}
func (s *RBACServiceImpl) DeleteRole(ctx context.Context, namespace, name string) error {
	if namespace != "" {
		err := global.KubeConfigSet.RbacV1().Roles(namespace).Delete(ctx, name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	} else {
		err := global.KubeConfigSet.RbacV1().ClusterRoles().Delete(ctx, name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *RBACServiceImpl) GetRoleBindingList(ctx context.Context, namespace, keyword string) ([]*rbac.RoleBindingResponse, error) {
	res := make([]*rbac.RoleBindingResponse, 0)
	if namespace != "" {
		list, err := global.KubeConfigSet.RbacV1().RoleBindings(namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}
		for _, item := range list.Items {
			if !strings.Contains(item.Name, keyword) {
				continue
			}
			res = append(res, &rbac.RoleBindingResponse{
				ResponseBase: &rbac.ResponseBase{
					Name:      item.Name,
					Namespace: item.Namespace,
					Age:       item.CreationTimestamp.Unix(),
				},
			})
		}
	} else {
		list, err := global.KubeConfigSet.RbacV1().ClusterRoleBindings().List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, err
		}

		for _, item := range list.Items {
			if !strings.Contains(item.Name, keyword) {
				continue
			}
			res = append(res, &rbac.RoleBindingResponse{
				ResponseBase: &rbac.ResponseBase{
					Name:      item.Name,
					Namespace: item.Namespace,
					Age:       item.CreationTimestamp.Unix(),
				},
			})
		}

	}
	return res, nil
}
func (s *RBACServiceImpl) GetRoleBindingDetail(ctx context.Context, namespace, name string) (*rbac.RoleBindingRequest, error) {
	res := rbac.NewRoleBindingRequest()
	if namespace != "" {
		rb, err := global.KubeConfigSet.RbacV1().RoleBindings(namespace).Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = rb.Name
		res.Namespace = rb.Namespace
		res.Labels = util.ListConverter(rb.Labels)
		res.Subjects = func(sbjs []rbacv1.Subject) []*rbac.ServiceAccountRequest {
			shouldReturn := make([]*rbac.ServiceAccountRequest, 0)
			for _, sbj := range sbjs {
				shouldReturn = append(shouldReturn, &rbac.ServiceAccountRequest{
					RequestBase: &rbac.RequestBase{
						Name:      sbj.Name,
						Namespace: sbj.Namespace,
					},
				})
			}
			return shouldReturn
		}(rb.Subjects)
		res.RoleRef = rb.RoleRef.Name
	} else {
		rb, err := global.KubeConfigSet.RbacV1().ClusterRoleBindings().Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = rb.Name
		res.Namespace = rb.Namespace
		res.Labels = util.ListConverter(rb.Labels)
		res.Subjects = func(sbjs []rbacv1.Subject) []*rbac.ServiceAccountRequest {
			shouldReturn := make([]*rbac.ServiceAccountRequest, 0)
			for _, sbj := range sbjs {
				shouldReturn = append(shouldReturn, &rbac.ServiceAccountRequest{
					RequestBase: &rbac.RequestBase{
						Name:      sbj.Name,
						Namespace: sbj.Namespace,
					},
				})
			}
			return shouldReturn
		}(rb.Subjects)
		res.RoleRef = rb.RoleRef.Name
	}
	return res, nil
}
func (s *RBACServiceImpl) CreateRoleBinding(ctx context.Context, req *rbac.RoleBindingRequest) (*rbac.RoleBindingRequest, error) {
	res := rbac.NewRoleBindingRequest()
	if req.Namespace != "" {
		rbK8S := &rbacv1.RoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name,
				Namespace: req.Namespace,
				Labels:    util.MapConverter(req.Labels),
			},
			Subjects: func(reqSbjs []*rbac.ServiceAccountRequest) []rbacv1.Subject {
				shouldReturn := make([]rbacv1.Subject, 0)
				for _, reqSbj := range reqSbjs {
					shouldReturn = append(shouldReturn, rbacv1.Subject{
						Name:      reqSbj.Name,
						Namespace: reqSbj.Namespace,
						Kind:      "User",
					})
				}
				return shouldReturn
			}(req.Subjects),
			RoleRef: rbacv1.RoleRef{
				Name:     req.RoleRef,
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "Role",
			},
		}
		newRole, err := global.KubeConfigSet.RbacV1().RoleBindings(req.Namespace).Create(ctx, rbK8S, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = newRole.Name
		res.Namespace = newRole.Namespace
		res.Labels = util.ListConverter(newRole.Labels)
		res.Subjects = func(sbjs []rbacv1.Subject) []*rbac.ServiceAccountRequest {
			shouldReturn := make([]*rbac.ServiceAccountRequest, 0)
			for _, sbj := range sbjs {
				shouldReturn = append(shouldReturn, &rbac.ServiceAccountRequest{
					RequestBase: &rbac.RequestBase{
						Name:      sbj.Name,
						Namespace: sbj.Namespace,
					},
				})
			}
			return shouldReturn
		}(newRole.Subjects)
		res.RoleRef = newRole.RoleRef.Name

	} else {
		rbCluster := &rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.Name,
				Namespace: req.Namespace,
				Labels:    util.MapConverter(req.Labels),
			},
			Subjects: func(reqSbjs []*rbac.ServiceAccountRequest) []rbacv1.Subject {
				shouldReturn := make([]rbacv1.Subject, 0)
				for _, reqSbj := range reqSbjs {
					shouldReturn = append(shouldReturn, rbacv1.Subject{
						Name:      reqSbj.Name,
						Namespace: reqSbj.Namespace,
						Kind:      "User",
					})
				}
				return shouldReturn
			}(req.Subjects),
			RoleRef: rbacv1.RoleRef{
				Name:     req.RoleRef,
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     "ClusterRole",
			},
		}
		newrb, err := global.KubeConfigSet.RbacV1().ClusterRoleBindings().Create(ctx, rbCluster, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		res.Name = newrb.Name
		res.Namespace = newrb.Namespace
		res.Labels = util.ListConverter(newrb.Labels)
		res.Subjects = func(sbjs []rbacv1.Subject) []*rbac.ServiceAccountRequest {
			shouldReturn := make([]*rbac.ServiceAccountRequest, 0)
			for _, sbj := range sbjs {
				shouldReturn = append(shouldReturn, &rbac.ServiceAccountRequest{
					RequestBase: &rbac.RequestBase{
						Name:      sbj.Name,
						Namespace: sbj.Namespace,
					},
				})
			}
			return shouldReturn
		}(newrb.Subjects)
		res.RoleRef = newrb.RoleRef.Name

	}
	return res, nil
}

func (s *RBACServiceImpl) DeleteRoleBinding(ctx context.Context, namespace, name string) error {

	if namespace == "" {
		err := global.KubeConfigSet.RbacV1().ClusterRoleBindings().Delete(ctx, name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	} else {
		err := global.KubeConfigSet.RbacV1().RoleBindings(namespace).Delete(ctx, name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
