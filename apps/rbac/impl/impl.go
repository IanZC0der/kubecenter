package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/rbac"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
