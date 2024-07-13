package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/secret"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func init() {
	ioc.DefaultControllerContainer().Register(&SecretServiceImpl{})
}

var _ secret.Service = &SecretServiceImpl{}

type SecretServiceImpl struct{}

func (s *SecretServiceImpl) Init() error {
	return nil
}

func (s *SecretServiceImpl) Name() string {
	return secret.AppName
}

func (s *SecretServiceImpl) GetSecretsList(ctx context.Context, namespace string, keyword string) ([]*secret.SecretResponse, error) {
	list, err := global.KubeConfigSet.CoreV1().Secrets(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	secretResList := make([]*secret.SecretResponse, 0)
	for _, item := range list.Items {
		if !strings.Contains(item.Name, keyword) {
			continue
		}
		secretRes := secret.GetSecretResponseFromK8SSecret(&item)
		secretResList = append(secretResList, secretRes)
	}
	return secretResList, err
}
func (s *SecretServiceImpl) GetSecretDetail(ctx context.Context, namespace string, name string) (*secret.SecretResponse, error) {
	secretK8s, err := global.KubeConfigSet.CoreV1().Secrets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	secretRes := secret.GetSecretResponseFromK8SSecret(secretK8s)
	return secretRes, err
}

func (s *SecretServiceImpl) DeleteSecret(ctx context.Context, namespace string, name string) error {
	return global.KubeConfigSet.CoreV1().Secrets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
func (s *SecretServiceImpl) UpdateSecret(ctx context.Context, req *secret.Secret) (*corev1.Secret, string, error) {
	k8sSecret := secret.GetK8SSecretFromSecretReq(req)

	_, err := global.KubeConfigSet.CoreV1().Secrets(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	if err != nil {
		//create
		sc, err := global.KubeConfigSet.CoreV1().Secrets(req.Namespace).Create(ctx, k8sSecret, metav1.CreateOptions{})
		if err != nil {
			return nil, "", err
		}
		return sc, "create secret success", nil
	}
	//update
	sc, err := global.KubeConfigSet.CoreV1().Secrets(req.Namespace).Update(ctx, k8sSecret, metav1.UpdateOptions{})
	if err != nil {
		return nil, "", err
	}
	return sc, "update secret success", nil

}
