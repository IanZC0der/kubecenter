package impl

import (
	"context"
	"github.com/IanZC0der/kubecenter/apps/nodes"
	"github.com/IanZC0der/kubecenter/global"
	"github.com/IanZC0der/kubecenter/ioc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func init() {
	ioc.DefaultControllerContainer().Register(&NodesServiceImpl{})
}

var _ nodes.Service = &NodesServiceImpl{}

type NodesServiceImpl struct{}

func (s *NodesServiceImpl) Init() error {
	return nil
}

func (s *NodesServiceImpl) Name() string {
	return nodes.AppName
}

func (s *NodesServiceImpl) GetNodeList(ctx context.Context, keyword string) ([]*nodes.Node, error) {
	c := context.TODO()

	list, err := global.KubeConfigSet.CoreV1().Nodes().List(c, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	nodesList := make([]*nodes.Node, 0)

	for _, node := range list.Items {
		if strings.Contains(node.Name, keyword) {
			nodesList = append(nodesList, nodes.GetNodeInfoFromK8SNode(&node))
		}
	}
	return nodesList, nil
}

func (s *NodesServiceImpl) GetNodeDetail(ctx context.Context, name string) (*nodes.Node, error) {
	c := context.TODO()
	k8sNode, err := global.KubeConfigSet.CoreV1().Nodes().Get(c, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	res := nodes.GetNodeDetailFromK8SNode(k8sNode)
	return res, nil
}
