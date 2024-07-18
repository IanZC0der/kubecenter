package impl

import (
	"context"
	"fmt"
	"github.com/IanZC0der/kubecenter/apps/k8sservice"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	"github.com/IanZC0der/kubecenter/util"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"strings"
)

type K8sServiceImpl struct{}

var _ k8sservice.Service = &K8sServiceImpl{}

func init() {
	ioc.DefaultControllerContainer().Register(&K8sServiceImpl{})
}

func (k *K8sServiceImpl) Init() error {
	return nil
}

func (k *K8sServiceImpl) Name() string {
	return k8sservice.AppName
}

func (k *K8sServiceImpl) GetSvcList(ctx context.Context, namespace string, keyword string) ([]*corev1.Service, error) {
	list, err := global.KubeConfigSet.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	res := make([]*corev1.Service, 0)
	for _, svc := range list.Items {
		if !strings.Contains(svc.Name, keyword) {
			continue
		}
		res = append(res, &svc)
	}
	return res, nil
}
func (k *K8sServiceImpl) GetSvcDetail(ctx context.Context, namespace string, name string) (*corev1.Service, error) {
	res, err := global.KubeConfigSet.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (k *K8sServiceImpl) CreateOrUpdateSvc(ctx context.Context, req *k8sservice.K8SService) (*corev1.Service, string, error) {
	svcPorts := make([]corev1.ServicePort, 0)
	for _, port := range req.Ports {
		svcPorts = append(svcPorts, corev1.ServicePort{
			Name: port.Name,
			Port: port.Port,
			TargetPort: intstr.IntOrString{
				Type:   intstr.Int,
				IntVal: port.TargetPort,
			},
			NodePort: port.NodePort,
		})
	}
	newSvc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      req.Name,
			Namespace: req.Namespace,
			Labels:    util.MapConverter(req.Labels),
		},
		Spec: corev1.ServiceSpec{
			Type:     req.Type,
			Selector: util.MapConverter(req.Selector),
			Ports:    svcPorts,
		},
	}

	existingSvc, err := global.KubeConfigSet.CoreV1().Services(req.Namespace).Get(ctx, req.Name, metav1.GetOptions{})
	var res *corev1.Service
	var msg string
	if err == nil {
		existingSvc.Spec = newSvc.Spec
		res, err = global.KubeConfigSet.CoreV1().Services(req.Namespace).Update(ctx, existingSvc, metav1.UpdateOptions{})
		if err != nil {
			msg = fmt.Sprintf("update service failed: %s", err.Error())
			return nil, msg, err
		} else {
			msg = fmt.Sprintf("update service succeed: %s", res.Name)
		}
	} else {
		res, err = global.KubeConfigSet.CoreV1().Services(req.Namespace).Create(ctx, newSvc, metav1.CreateOptions{})
		if err != nil {
			msg = fmt.Sprintf("create service failed: %s", err.Error())
			return nil, msg, err
		} else {
			msg = fmt.Sprintf("create service succeed: %s", res.Name)
		}
	}
	return res, msg, nil
}

func (k *K8sServiceImpl) DeleteSvc(ctx context.Context, namespace, name string) error {
	return global.KubeConfigSet.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}
