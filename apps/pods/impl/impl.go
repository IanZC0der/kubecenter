package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/kube-openapi/pkg/util"
	"strings"

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

func (s *PodsServerImpl) GetNamespaceList(ctx context.Context) (*pods.NamespaceList, error) {
	namespaceList := pods.NewNamespaceList()
	c := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Namespaces().List(c, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}
	for _, namespace := range list.Items {
		namespaceList.Items = append(namespaceList.Items, &pods.Namespace{
			Name:              namespace.Name,
			CreationTimestamp: namespace.CreationTimestamp.Unix(),
			Status:            string(namespace.Status.Phase),
		})
	}
	return namespaceList, nil
}

func (s *PodsServerImpl) GetPodsListUnderNamespaceWithKeyword(ctx context.Context, namespace string, keyword string) (*pods.PodsList, error) {
	podsList := pods.NewPodsItemsList()

	c := context.TODO()
	list, err := global.KubeConfigSet.CoreV1().Pods(namespace).List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, item := range list.Items {
		if strings.Contains(item.Name, keyword) {
			podsList.Items = append(podsList.Items, pods.GetPodListItemFromPod(&item))
		}
	}
	return podsList, nil
}
func (s *PodsServerImpl) Init() error {
	return nil
}

func (s *PodsServerImpl) Name() string {
	return pods.AppName
}
