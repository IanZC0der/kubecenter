package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/IanZC0der/kubecenter/apps/pods"
)

func init() {
	ioc.DefaultControllerContainer().Register(&PodsServerImpl{})
}

var _ pods.Service = &PodsServerImpl{}

type PodsServerImpl struct {
}

func (s *PodsServerImpl) GetPods(ctx context.Context) (*pods.Pods, error) {
	podsList := pods.NewPodsList()
	c := context.TODO()
	pods, err := global.KubeConfigSet.CoreV1().Pods("").List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		podsList.Items[pod.Namespace] = append(podsList.Items[pod.Namespace], pod.Name)
	}

	return podsList, nil

}
func (s *PodsServerImpl) Init() error {
	return nil
}

func (s *PodsServerImpl) Name() string {
	return pods.AppName
}
